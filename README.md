# go-editor

By passing bytes you can edit with editor
* Pass the most favorite editor as an argument.

## sample
```
	var data []byte
	data = []byte("For example, stringdata hogehogehoge")

	ret, err := editor.Call("vim",data) //[]byte, error (string, []byte)
	if err != nil {
		return
	}
	fmt.Printf("%s",ret)
```

## Other's

* Although there is no relation. For the Python, write like this.
```
def callEditor(editor="vim", str=""):
    buf=str
    EDITOR = os.environ.get('EDITOR',editor)
    with tempfile.NamedTemporaryFile(suffix=".tmp") as tf:
        tf.write(buf.encode())
        tf.flush()
        call([EDITOR, tf.name])
        tf.seek(0)
        buf = tf.read()
    return buf.decode()
```
