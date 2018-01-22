# TODO
さっぱりわかりません。
CountingWriter(w io.Writer) (io.Writer, *int64) 
の返り値をreturn &result, &(result.counter)
とすると、counterは更新されるのは理解できますが、 この時の書き込んだ元の値io.writerは消えるのですが良いのですか？
毎回、CountingWriterを更新すると、counterが更新されるという認識で正しいですか？