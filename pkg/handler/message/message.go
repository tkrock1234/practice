package message

func (m *Message) func NewMessage(c context.Context) (*datastore.Key, error){
	key := datastore.NewInCompleteKey(c, "Message", nil)
	return datastore.Put(c, key, m)
}
