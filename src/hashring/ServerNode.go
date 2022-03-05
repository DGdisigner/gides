package hashring

type ServerNode struct {
	Host       string
	Port       string
	ServerHash uint32
	Next       *ServerNode
	First      *ServerNode
}

func (s *ServerNode) ring(keyHash uint32) *ServerNode {
	if s.ServerHash > keyHash {
		return s
	} else {
		if s.Next == nil {
			return s.First
		} else {
			return s.Next.ring(keyHash)
		}
	}
}

func (s *ServerNode) Find(key string) {
	//keyHash := crc32.ChecksumIEEE([]byte(key))
	//server := s.First.ring(keyHash)

}
