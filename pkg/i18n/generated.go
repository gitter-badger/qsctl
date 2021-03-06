package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// initEnUS will init en_US support.
func initEnUS(tag language.Tag) {
	_ = message.SetString(tag, "%s  %s  %s  %s\n", "%s  %s  %s  %s\n")
	_ = message.SetString(tag, "-r is required to delete a directory", "-r is required to delete a directory")
	_ = message.SetString(tag, "AccessKey and SecretKey not found. Please setup your config now, or exit and setup manually.", "AccessKey and SecretKey not found. Please setup your config now, or exit and setup manually.")
	_ = message.SetString(tag, "Bucket <%s> created.\n", "Bucket <%s> created.\n")
	_ = message.SetString(tag, "Bucket <%s> removed.\n", "Bucket <%s> removed.\n")
	_ = message.SetString(tag, "Cat object: qsctl cat qs://prefix/a", "Cat object: qsctl cat qs://prefix/a")
	_ = message.SetString(tag, "Config not loaded, use default and environment value instead.", "Config not loaded, use default and environment value instead.")
	_ = message.SetString(tag, "Copy file: qsctl cp /path/to/file qs://prefix/a", "Copy file: qsctl cp /path/to/file qs://prefix/a")
	_ = message.SetString(tag, "Copy folder: qsctl cp qs://prefix/a /path/to/folder -r", "Copy folder: qsctl cp qs://prefix/a /path/to/folder -r")
	_ = message.SetString(tag, "Delete an empty qingstor bucket or forcely delete nonempty qingstor bucket.", "Delete an empty qingstor bucket or forcely delete nonempty qingstor bucket.")
	_ = message.SetString(tag, "Dir <%s> and <%s> synced.\n", "Dir <%s> and <%s> synced.\n")
	_ = message.SetString(tag, "Dir <%s> copied to <%s>.\n", "Dir <%s> copied to <%s>.\n")
	_ = message.SetString(tag, "Dir <%s> moved to <%s>.\n", "Dir <%s> moved to <%s>.\n")
	_ = message.SetString(tag, "Dir <%s> removed.\n", "Dir <%s> removed.\n")
	_ = message.SetString(tag, "File <%s> copied to <%s>.\n", "File <%s> copied to <%s>.\n")
	_ = message.SetString(tag, "File <%s> moved to <%s>.\n", "File <%s> moved to <%s>.\n")
	_ = message.SetString(tag, "File <%s> removed.\n", "File <%s> removed.\n")
	_ = message.SetString(tag, "Key: %s", "Key: %s")
	_ = message.SetString(tag, "List bucket's all objects: qsctl ls qs://bucket-name -R", "List bucket's all objects: qsctl ls qs://bucket-name -R")
	_ = message.SetString(tag, "List buckets: qsctl ls", "List buckets: qsctl ls")
	_ = message.SetString(tag, "List objects by long format: qsctl ls qs://bucket-name -l", "List objects by long format: qsctl ls qs://bucket-name -l")
	_ = message.SetString(tag, "List objects with prefix: qsctl ls qs://bucket-name/prefix", "List objects with prefix: qsctl ls qs://bucket-name/prefix")
	_ = message.SetString(tag, "Load config failed [%v]", "Load config failed [%v]")
	_ = message.SetString(tag, "MD5: %s", "MD5: %s")
	_ = message.SetString(tag, "Make bucket: qsctl mb bucket-name", "Make bucket: qsctl mb bucket-name")
	_ = message.SetString(tag, "Move file: qsctl mv /path/to/file qs://prefix/a", "Move file: qsctl mv /path/to/file qs://prefix/a")
	_ = message.SetString(tag, "Move folder: qsctl mv qs://prefix/a /path/to/folder -r", "Move folder: qsctl mv qs://prefix/a /path/to/folder -r")
	_ = message.SetString(tag, "Presign object: qsctl qs://bucket-name/object-name", "Presign object: qsctl qs://bucket-name/object-name")
	_ = message.SetString(tag, "Read from stdin: cat /path/to/file | qsctl cp - qs://prefix/stdin", "Read from stdin: cat /path/to/file | qsctl cp - qs://prefix/stdin")
	_ = message.SetString(tag, "Remove a single object: qsctl rm qs://bucket-name/object-key", "Remove a single object: qsctl rm qs://bucket-name/object-key")
	_ = message.SetString(tag, "Size: %s", "Size: %s")
	_ = message.SetString(tag, "Stat object: qsctl stat qs://prefix/a", "Stat object: qsctl stat qs://prefix/a")
	_ = message.SetString(tag, "StorageClass: %s", "StorageClass: %s")
	_ = message.SetString(tag, "Sync QS-Directory to local directory: qsctl sync qs://bucket-name/test/ test_local/", "Sync QS-Directory to local directory: qsctl sync qs://bucket-name/test/ test_local/")
	_ = message.SetString(tag, "Sync local directory to QS-Directory: qsctl sync . qs://bucket-name", "Sync local directory to QS-Directory: qsctl sync . qs://bucket-name")
	_ = message.SetString(tag, "Sync skip updating files that already exist on receiver: qsctl sync . qs://bucket-name --ignore-existing", "Sync skip updating files that already exist on receiver: qsctl sync . qs://bucket-name --ignore-existing")
	_ = message.SetString(tag, "Tee object: qsctl tee qs://prefix/a", "Tee object: qsctl tee qs://prefix/a")
	_ = message.SetString(tag, "Type: %s", "Type: %s")
	_ = message.SetString(tag, "UpdatedAt: %s", "UpdatedAt: %s")
	_ = message.SetString(tag, "Write to stdout: qsctl cp qs://prefix/b - > /path/to/file", "Write to stdout: qsctl cp qs://prefix/b - > /path/to/file")
	_ = message.SetString(tag, "Your config has been set to <%v>. You can still modify it manually.", "Your config has been set to <%v>. You can still modify it manually.")
	_ = message.SetString(tag, "assign config path manually", "assign config path manually")
	_ = message.SetString(tag, "cat a remote object to stdout", "cat a remote object to stdout")
	_ = message.SetString(tag, "copy directory recursively", "copy directory recursively")
	_ = message.SetString(tag, "copy from/to qingstor", "copy from/to qingstor")
	_ = message.SetString(tag, "delete a bucket", "delete a bucket")
	_ = message.SetString(tag, "delete an empty bucket: qsctl rb qs://bucket-name", "delete an empty bucket: qsctl rb qs://bucket-name")
	_ = message.SetString(tag, "enable benchmark or not", "enable benchmark or not")
	_ = message.SetString(tag, "forcely delete a nonempty bucket: qsctl rb qs://bucket-name -f", "forcely delete a nonempty bucket: qsctl rb qs://bucket-name -f")
	_ = message.SetString(tag, "get the pre-signed URL by the object key", "get the pre-signed URL by the object key")
	_ = message.SetString(tag, "help for this command", "help for this command")
	_ = message.SetString(tag, "in which zone to do the operation", "in which zone to do the operation")
	_ = message.SetString(tag, "in which zone to make the bucket (required)", "in which zone to make the bucket (required)")
	_ = message.SetString(tag, "list objects or buckets", "list objects or buckets")
	_ = message.SetString(tag, "make a new bucket", "make a new bucket")
	_ = message.SetString(tag, "move directory recursively", "move directory recursively")
	_ = message.SetString(tag, "move from/to qingstor", "move from/to qingstor")
	_ = message.SetString(tag, "print logs for debug", "print logs for debug")
	_ = message.SetString(tag, "qsctl cat can cat a remote object to stdout", "qsctl cat can cat a remote object to stdout")
	_ = message.SetString(tag, "qsctl cp can copy file/folder/stdin to qingstor or copy qingstor objects to local/stdout", "qsctl cp can copy file/folder/stdin to qingstor or copy qingstor objects to local/stdout")
	_ = message.SetString(tag, "qsctl mv can move file/folder to qingstor or move qingstor objects to local", "qsctl mv can move file/folder to qingstor or move qingstor objects to local")
	_ = message.SetString(tag, "qsctl rb delete a qingstor bucket", "qsctl rb delete a qingstor bucket")
	_ = message.SetString(tag, "qsctl rm remove the object with given object key", "qsctl rm remove the object with given object key")
	_ = message.SetString(tag, "qsctl stat show the detailed info of this object", "qsctl stat show the detailed info of this object")
	_ = message.SetString(tag, "recursively delete keys under a specific prefix", "recursively delete keys under a specific prefix")
	_ = message.SetString(tag, "recursively list subdirectories encountered", "recursively list subdirectories encountered")
	_ = message.SetString(tag, "remove a remote object", "remove a remote object")
	_ = message.SetString(tag, "stat a remote object", "stat a remote object")
	_ = message.SetString(tag, "sync between local directory and QS-Directory", "sync between local directory and QS-Directory")
	_ = message.SetString(tag, "tee a remote object from stdin", "tee a remote object from stdin")
	_ = message.SetString(tag, "the number of seconds until the pre-signed URL expires. Default is 300 seconds", "the number of seconds until the pre-signed URL expires. Default is 300 seconds")
	_ = message.SetString(tag, `expected size of the input file
accept: 100MB, 1.8G
(only used and required for input from stdin)`, `expected size of the input file
accept: 100MB, 1.8G
(only used and required for input from stdin)`)
	_ = message.SetString(tag, `maximum content loaded in memory
(only used for input from stdin)`, `maximum content loaded in memory
(only used for input from stdin)`)
	_ = message.SetString(tag, `qsctl ls can list all qingstor buckets or qingstor keys under a prefix.`, `qsctl ls can list all qingstor buckets or qingstor keys under a prefix.`)
	_ = message.SetString(tag, `qsctl mb can make a new bucket with the specific name,

bucket name should follow DNS name rule with:
* length between 6 and 63;
* can only contains lowercase letters, numbers and hyphen -
* must start and end with lowercase letter or number
* must not be an available IP address
	`, `qsctl mb can make a new bucket with the specific name,

bucket name should follow DNS name rule with:
* length between 6 and 63;
* can only contains lowercase letters, numbers and hyphen -
* must start and end with lowercase letter or number
* must not be an available IP address
	`)
	_ = message.SetString(tag, `qsctl presign can generate a pre-signed URL for the object.
Within the given expire time, anyone who receives this URL can retrieve
the object with an HTTP GET request. If an object belongs to a public bucket,
generate a URL spliced by bucket name, zone and its name, anyone who receives
this URL can always retrieve the object with an HTTP GET request.`, `qsctl presign can generate a pre-signed URL for the object.
Within the given expire time, anyone who receives this URL can retrieve
the object with an HTTP GET request. If an object belongs to a public bucket,
generate a URL spliced by bucket name, zone and its name, anyone who receives
this URL can always retrieve the object with an HTTP GET request.`)
	_ = message.SetString(tag, `qsctl sync between local directory and QS-Directory. The first path argument
is the source directory and second the destination directory.

When a key(file) already exists in the destination directory, program will compare
the modified time of source file(key) and destination key(file). The destination
key(file) will be overwritten only if the source one newer than destination one.`, `qsctl sync between local directory and QS-Directory. The first path argument
is the source directory and second the destination directory.

When a key(file) already exists in the destination directory, program will compare
the modified time of source file(key) and destination key(file). The destination
key(file) will be overwritten only if the source one newer than destination one.`)
	_ = message.SetString(tag, `qsctl tee can tee a remote object from stdin.

NOTICE: qsctl will not tee the content to stdout like linux tee command does.
`, `qsctl tee can tee a remote object from stdin.

NOTICE: qsctl will not tee the content to stdout like linux tee command does.
`)
	_ = message.SetString(tag, `skip creating new files in dest dirs, only copy newer by time`, `skip creating new files in dest dirs, only copy newer by time`)
	_ = message.SetString(tag, `use the specified FORMAT instead of the default;
output a newline after each use of FORMAT

The valid format sequences for files:

  %F   file type
  %h   content md5 of the file
  %n   file name
  %s   total size, in bytes
  %y   time of last data modification, human-readable, e.g: 2006-01-02 15:04:05 +0000 UTC
  %Y   time of last data modification, seconds since Epoch
	`, `use the specified FORMAT instead of the default;
output a newline after each use of FORMAT

The valid format sequences for files:

  %F   file type
  %h   content md5 of the file
  %n   file name
  %s   total size, in bytes
  %y   time of last data modification, human-readable, e.g: 2006-01-02 15:04:05 +0000 UTC
  %Y   time of last data modification, seconds since Epoch
	`)
	_ = message.SetString(tag, `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`, `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`)
	_ = message.SetString(tag, `{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}`, `{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}`)
}

// initZhCN will init zh_CN support.
func initZhCN(tag language.Tag) {
	_ = message.SetString(tag, "%s  %s  %s  %s\n", "%s  %s  %s  %s\n")
	_ = message.SetString(tag, "-r is required to delete a directory", "删除目录必须要有 -r 参数")
	_ = message.SetString(tag, "AccessKey and SecretKey not found. Please setup your config now, or exit and setup manually.", "AccessKey 和 SecretKey 未找到。现在请设置您的配置，或者退出以手动设置。")
	_ = message.SetString(tag, "Bucket <%s> created.\n", "Bucket <%s> 已创建。\n")
	_ = message.SetString(tag, "Bucket <%s> removed.\n", "Bucket <%s> 已删除。\n")
	_ = message.SetString(tag, "Cat object: qsctl cat qs://prefix/a", "输出一个文件的内容到标准输出: qsctl cat qs://prefix/a")
	_ = message.SetString(tag, "Config not loaded, use default and environment value instead.", "配置未加载，使用默认值和环境变量代替。")
	_ = message.SetString(tag, "Copy file: qsctl cp /path/to/file qs://prefix/a", "复制文件: qsctl cp / path/to/file qs://prefix/a")
	_ = message.SetString(tag, "Copy folder: qsctl cp qs://prefix/a /path/to/folder -r", "复制文件夹: qsctl cp qs://prefix/a /path/to/folder -r")
	_ = message.SetString(tag, "Delete an empty qingstor bucket or forcely delete nonempty qingstor bucket.", "删除空 Bucket 或强制删除非空 Bucket。")
	_ = message.SetString(tag, "Dir <%s> and <%s> synced.\n", "文件夹 <%s> and <%s> 已同步。\n")
	_ = message.SetString(tag, "Dir <%s> copied to <%s>.\n", "文件夹 <%s> 已复制到 <%s>.\n")
	_ = message.SetString(tag, "Dir <%s> moved to <%s>.\n", "文件夹 <%s> 已移动到 <%s>.\n")
	_ = message.SetString(tag, "Dir <%s> removed.\n", "文件夹 <%s> 已删除。\n")
	_ = message.SetString(tag, "File <%s> copied to <%s>.\n", "文件 <%s> 已复制到 <%s>.\n")
	_ = message.SetString(tag, "File <%s> moved to <%s>.\n", "文件 <%s> 已移动到 <%s>.\n")
	_ = message.SetString(tag, "File <%s> removed.\n", "文件 <%s> 已删除。\n")
	_ = message.SetString(tag, "Key: %s", "名称: %s")
	_ = message.SetString(tag, "List bucket's all objects: qsctl ls qs://bucket-name -R", "列出 Bucket 中的所有对象: qsctl ls qs://bucket-name -R")
	_ = message.SetString(tag, "List buckets: qsctl ls", "列出 Bucket: qsctl ls")
	_ = message.SetString(tag, "List objects by long format: qsctl ls qs://bucket-name -l", "使用详细格式列出对象: qsctl ls qs://bucket-name -l")
	_ = message.SetString(tag, "List objects with prefix: qsctl ls qs://bucket-name/prefix", "列出带指定前缀的对象: qsctl ls qs://bucket-name/prefix")
	_ = message.SetString(tag, "Load config failed [%v]", "加载配置失败 [%v]")
	_ = message.SetString(tag, "MD5: %s", "MD5: %s")
	_ = message.SetString(tag, "Make bucket: qsctl mb bucket-name", "创建一个 Bucket: qsctl mb bucket-name")
	_ = message.SetString(tag, "Move file: qsctl mv /path/to/file qs://prefix/a", "移动文件: qsctl mv /path/to/file qs://prefix/a")
	_ = message.SetString(tag, "Move folder: qsctl mv qs://prefix/a /path/to/folder -r", "移动文件夹: qsctl mv qs://prefix/a /path/to/folder -r")
	_ = message.SetString(tag, "Presign object: qsctl qs://bucket-name/object-name", "预签名对象: qsctl qs://bucket-name/object-name")
	_ = message.SetString(tag, "Read from stdin: cat /path/to/file | qsctl cp - qs://prefix/stdin", "从 stdin 读取并上传: cat /path/to/file | qsctl cp - qs://prefix/stdin")
	_ = message.SetString(tag, "Remove a single object: qsctl rm qs://bucket-name/object-key", "删除单个对象: qsctl rm qs://bucket-name/object-key")
	_ = message.SetString(tag, "Size: %s", "大小: %s")
	_ = message.SetString(tag, "Stat object: qsctl stat qs://prefix/a", "查看文件信息: qsctl stat qs://prefix/a")
	_ = message.SetString(tag, "StorageClass: %s", "存储类型: %s")
	_ = message.SetString(tag, "Sync QS-Directory to local directory: qsctl sync qs://bucket-name/test/ test_local/", "同步 QS-Directory 到本地目录: qsctl sync qs://bucket-name/test/ test_local/")
	_ = message.SetString(tag, "Sync local directory to QS-Directory: qsctl sync . qs://bucket-name", "同步本地目录到 QS-Directory: qsctl sync . qs://bucket-name")
	_ = message.SetString(tag, "Sync skip updating files that already exist on receiver: qsctl sync . qs://bucket-name --ignore-existing", "同步并跳过更新接收端已经存在的文件: qsctl sync . qs://bucket-name --ignore-existing")
	_ = message.SetString(tag, "Tee object: qsctl tee qs://prefix/a", "输出一个文件的内容到标准输出: qsctl tee qs://prefix/a")
	_ = message.SetString(tag, "Type: %s", "类型: %s")
	_ = message.SetString(tag, "UpdatedAt: %s", "更新于: %s")
	_ = message.SetString(tag, "Write to stdout: qsctl cp qs://prefix/b - > /path/to/file", "写入到标准输出: qsctl cp qs://prefix/b - > /path/to/file")
	_ = message.SetString(tag, "Your config has been set to <%v>. You can still modify it manually.", "您的配置已设置为 <%v>。您仍然可以手动修改它。")
	_ = message.SetString(tag, "assign config path manually", "手动分配配置路径")
	_ = message.SetString(tag, "cat a remote object to stdout", "输出远程对象内容到标准输出")
	_ = message.SetString(tag, "copy directory recursively", "递归复制目录")
	_ = message.SetString(tag, "copy from/to qingstor", "复制从/到 QingStor 对象存储")
	_ = message.SetString(tag, "delete a bucket", "删除一个 Bucket")
	_ = message.SetString(tag, "delete an empty bucket: qsctl rb qs://bucket-name", "删除空 Bucket: qsctl rb qs://bucket-name")
	_ = message.SetString(tag, "enable benchmark or not", "启用性能测试与否")
	_ = message.SetString(tag, "forcely delete a nonempty bucket: qsctl rb qs://bucket-name -f", "强制删除一个非空桶: qsctl rb qs://bucket-name -f")
	_ = message.SetString(tag, "get the pre-signed URL by the object key", "通过对象键获取预签名的 URL")
	_ = message.SetString(tag, "help for this command", "帮助信息")
	_ = message.SetString(tag, "in which zone to do the operation", "在哪个区域执行操作")
	_ = message.SetString(tag, "in which zone to make the bucket (required)", "在哪个区域创建 Bucket (必须参数)")
	_ = message.SetString(tag, "list objects or buckets", "列出对象或 Bucket")
	_ = message.SetString(tag, "make a new bucket", "创建一个新的 Bucket")
	_ = message.SetString(tag, "move directory recursively", "递归移动目录")
	_ = message.SetString(tag, "move from/to qingstor", "移动从/到 QingStor 对象存储")
	_ = message.SetString(tag, "print logs for debug", "打印调试日志")
	_ = message.SetString(tag, "qsctl cat can cat a remote object to stdout", "qsctl cat 可以将远程对象内容输出到标准输出")
	_ = message.SetString(tag, "qsctl cp can copy file/folder/stdin to qingstor or copy qingstor objects to local/stdout", "qsctl cp 可以将文件/文件夹/stdin 复制到 QingStor 对象存储或复制对象到本地/stdout")
	_ = message.SetString(tag, "qsctl mv can move file/folder to qingstor or move qingstor objects to local", "qsctl mv 可以将文件/文件夹移动到 QingStor 对象存储或移动对象到本地")
	_ = message.SetString(tag, "qsctl rb delete a qingstor bucket", "qscl rb 将删除一个 Bucket")
	_ = message.SetString(tag, "qsctl rm remove the object with given object key", "qsctl rm 将删除给定 Object Key 的对象")
	_ = message.SetString(tag, "qsctl stat show the detailed info of this object", "qsctl stat 将显示此对象的详细信息")
	_ = message.SetString(tag, "recursively delete keys under a specific prefix", "递归删除指定前缀下的对象")
	_ = message.SetString(tag, "recursively list subdirectories encountered", "递归列出遇到的子目录")
	_ = message.SetString(tag, "remove a remote object", "删除远程对象")
	_ = message.SetString(tag, "stat a remote object", "查看远程对象的信息")
	_ = message.SetString(tag, "sync between local directory and QS-Directory", "同步本地目录和对象存储目录")
	_ = message.SetString(tag, "tee a remote object from stdin", "从标准输入读取内容并上传")
	_ = message.SetString(tag, "the number of seconds until the pre-signed URL expires. Default is 300 seconds", "预签名URL到期前的秒数。默认值为300秒")
	_ = message.SetString(tag, `expected size of the input file
accept: 100MB, 1.8G
(only used and required for input from stdin)`, `预计输入文件的大小
接受的大小形似: 100MB, 1.8G
(仅用于标准输入) `)
	_ = message.SetString(tag, `maximum content loaded in memory
(only used for input from stdin)`, `在内存中加载的最大内容
(仅用于标准输入)`)
	_ = message.SetString(tag, `qsctl ls can list all qingstor buckets or qingstor keys under a prefix.`, `qsctl ls 可以列出所有 Bucket 或者按前缀列出 QingStor 对象。`)
	_ = message.SetString(tag, `qsctl mb can make a new bucket with the specific name,

bucket name should follow DNS name rule with:
* length between 6 and 63;
* can only contains lowercase letters, numbers and hyphen -
* must start and end with lowercase letter or number
* must not be an available IP address
	`, `qsctl mb 可以用指定名称创建一个新的 Bucket。

bucket 名称应该遵循DNS名称规则:
* 长度介于 6 到 63 之间。
* 只能包含小写字母 数字和连线 -
* 必须以小写字母或数字开头和结尾
* 不能是可用的 IP 地址
`)
	_ = message.SetString(tag, `qsctl presign can generate a pre-signed URL for the object.
Within the given expire time, anyone who receives this URL can retrieve
the object with an HTTP GET request. If an object belongs to a public bucket,
generate a URL spliced by bucket name, zone and its name, anyone who receives
this URL can always retrieve the object with an HTTP GET request.`, `qsctl presign 可以为对象生成一个预签名的 URL。
在给定的时间内，任何拥有该链接的人都可以通过 HTTP GET 请求获取这个文件。如果这个文件属于一个公开的 Bucket，任何拥有该链接的人总是能够通过 HTTP GET 请求访问这个文件。`)
	_ = message.SetString(tag, `qsctl sync between local directory and QS-Directory. The first path argument
is the source directory and second the destination directory.

When a key(file) already exists in the destination directory, program will compare
the modified time of source file(key) and destination key(file). The destination
key(file) will be overwritten only if the source one newer than destination one.`, `qsctl 同步本地目录和 QS-Directory。第一个路径参数
是源目录，第二个目标目录。

当目标目录中已经存在一个密钥(文件) 时，程序将比较
源文件(密钥) 和目标密钥(文件) 的修改时间。 目标
密钥(文件) 只有当源比目标更新时才会被覆盖。`)
	_ = message.SetString(tag, `qsctl tee can tee a remote object from stdin.

NOTICE: qsctl will not tee the content to stdout like linux tee command does.
`, `qsctl tee 可以从 stdin 读取并上传文件。

注意: qsctl 将不会像 Linux tee 命令那样将内容绑定到标准输出。
`)
	_ = message.SetString(tag, `skip creating new files in dest dirs, only copy newer by time`, `跳过在目标目录中创建新文件，只能复制新创建的文件`)
	_ = message.SetString(tag, `use the specified FORMAT instead of the default;
output a newline after each use of FORMAT

The valid format sequences for files:

  %F   file type
  %h   content md5 of the file
  %n   file name
  %s   total size, in bytes
  %y   time of last data modification, human-readable, e.g: 2006-01-02 15:04:05 +0000 UTC
  %Y   time of last data modification, seconds since Epoch
	`, `使用指定的 FORMAT 而不是默认值；
每次使用 FORMAT

文件的有效格式序列后输出一条新行:

  %F   文件类型
  %h   文件的 md5 内容
  %n   文件名
  %s   总大小. 使用字节
  %y   最后一次数据修改的时间，比如: 2006-01-02 15:04:05+00UTC
  %Y   最后一次数据修改时间， Unix 时间戳
`)
	_ = message.SetString(tag, `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`, `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`)
	_ = message.SetString(tag, `{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}`, `{{with .Name}}{{printf "%s " .}}{{end}}{{printf "版本 %s" .Version}}`)
}
