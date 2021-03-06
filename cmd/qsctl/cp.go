package main

import (
	"fmt"
	"path/filepath"

	"github.com/Xuanwo/storage/types"
	"github.com/Xuanwo/storage/types/pairs"
	"github.com/qingstor/noah/task"
	"github.com/spf13/cobra"

	"github.com/qingstor/qsctl/v2/cmd/qsctl/taskutils"
	"github.com/qingstor/qsctl/v2/constants"
	"github.com/qingstor/qsctl/v2/pkg/i18n"
	"github.com/qingstor/qsctl/v2/utils"
)

var cpInput struct {
	ExpectSize           string
	MaximumMemoryContent string
	Recursive            bool
}

// CpCommand will handle copy command.
var CpCommand = &cobra.Command{
	Use:   "cp <source-path> <dest-path>",
	Short: i18n.Sprintf("copy from/to qingstor"),
	Long:  i18n.Sprintf("qsctl cp can copy file/folder/stdin to qingstor or copy qingstor objects to local/stdout"),
	Example: utils.AlignPrintWithColon(
		i18n.Sprintf("Copy file: qsctl cp /path/to/file qs://prefix/a"),
		i18n.Sprintf("Copy folder: qsctl cp qs://prefix/a /path/to/folder -r"),
		i18n.Sprintf("Read from stdin: cat /path/to/file | qsctl cp - qs://prefix/stdin"),
		i18n.Sprintf("Write to stdout: qsctl cp qs://prefix/b - > /path/to/file"),
	),
	Args: cobra.ExactArgs(2),
	RunE: cpRun,
}

func initCpFlag() {
	CpCommand.PersistentFlags().StringVar(&cpInput.ExpectSize,
		constants.ExpectSizeFlag,
		"",
		i18n.Sprintf(`expected size of the input file
accept: 100MB, 1.8G
(only used and required for input from stdin)`),
	)
	CpCommand.PersistentFlags().StringVar(&cpInput.MaximumMemoryContent,
		constants.MaximumMemoryContentFlag,
		"",
		i18n.Sprintf(`maximum content loaded in memory
(only used for input from stdin)`),
	)
	CpCommand.Flags().BoolVarP(&cpInput.Recursive,
		constants.RecursiveFlag,
		"r",
		false,
		i18n.Sprintf("copy directory recursively"),
	)
}

func cpRun(_ *cobra.Command, args []string) (err error) {
	rootTask := taskutils.NewBetweenStorageTask(10)
	err = utils.ParseBetweenStorageInput(rootTask, args[0], args[1])
	if err != nil {
		return
	}

	if rootTask.GetSourceType() == types.ObjectTypeDir && !cpInput.Recursive {
		return fmt.Errorf(i18n.Sprintf("-r is required to delete a directory"))
	}

	if err = HandleBetweenStorageWdAndPath(rootTask, cpInput.Recursive); err != nil {
		return err
	}

	if cpInput.Recursive {
		t := task.NewCopyDir(rootTask)
		t.SetCheckTasks(nil)
		t.Run()

		if t.GetFault().HasError() {
			return t.GetFault()
		}
		i18n.Printf("Dir <%s> copied to <%s>.\n", t.GetSourcePath(), t.GetDestinationPath())
		return nil
	}

	t := task.NewCopyFile(rootTask)
	t.SetCheckTasks(nil)
	t.Run()
	if t.GetFault().HasError() {
		return t.GetFault()
	}
	i18n.Printf("File <%s> copied to <%s>.\n", t.GetSourcePath(), t.GetDestinationPath())
	return
}

// HandleBetweenStorageWdAndPath set work dir and path for cp cmd.
func HandleBetweenStorageWdAndPath(t *taskutils.BetweenStorageTask, recursive bool) error {
	// In operation cp, we set source storage to dir of the source path.
	srcPath, err := filepath.Abs(t.GetSourcePath())
	if err != nil {
		return err
	}
	if err = t.GetSourceStorage().Init(pairs.WithWorkDir(filepath.Dir(srcPath))); err != nil {
		return err
	}
	t.SetSourcePath(filepath.Base(srcPath))

	// if source path == source work dir，which means cp from '/', transfer to cp /*
	if filepath.Base(srcPath) == filepath.Dir(srcPath) {
		t.SetSourcePath("")
	}

	// Destination path depends on different condition.
	dstPath, err := filepath.Abs(t.GetDestinationPath())
	if err != nil {
		return err
	}
	// if copy dir
	if recursive {
		if err := t.GetDestinationStorage().Init(pairs.WithWorkDir(dstPath)); err != nil {
			return err
		}
		t.SetDestinationPath("")
		return nil
	}
	// NOT copy dir. Copy file to a dir, we need to get destination key from the source.
	if t.GetDestinationType() == types.ObjectTypeDir {
		if err := t.GetDestinationStorage().Init(pairs.WithWorkDir(dstPath)); err != nil {
			return err
		}
		t.SetDestinationPath(t.GetSourcePath())
		return nil
	}
	// Copy to a file, get destination directly.
	if err := t.GetDestinationStorage().Init(pairs.WithWorkDir(filepath.Dir(dstPath))); err != nil {
		return err
	}
	t.SetDestinationPath(filepath.Base(dstPath))
	return nil
}
