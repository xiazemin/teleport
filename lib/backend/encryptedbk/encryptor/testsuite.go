// package test contains a backend acceptance test suite that is backend implementation independant
// each backend will use the suite to test itself
package encryptor

import (
	. "github.com/gravitational/teleport/Godeps/_workspace/src/gopkg.in/check.v1"
)

type encryptorSuite struct {
	E Encryptor
}

func (s *encryptorSuite) EncryptDecrypt(c *C) {
	s1 := "first string"
	s2 := "second"
	s3 := "third text"

	e1, err := s.E.Encrypt([]byte(s1))
	c.Assert(err, IsNil)
	e11, err := s.E.Encrypt([]byte(s1))
	c.Assert(err, IsNil)
	e2, err := s.E.Encrypt([]byte(s2))
	c.Assert(err, IsNil)
	e3, err := s.E.Encrypt([]byte(s3))
	c.Assert(err, IsNil)

	c.Assert(s1 == string(e1), Equals, false)
	c.Assert(s2 == string(e2), Equals, false)
	c.Assert(s3 == string(e3), Equals, false)

	c.Assert(string(e1) == string(e2), Equals, false)
	c.Assert(string(e1) == string(e3), Equals, false)
	c.Assert(string(e2) == string(e3), Equals, false)

	d1, err := s.E.Decrypt([]byte(e1))
	c.Assert(err, IsNil)
	d11, err := s.E.Decrypt([]byte(e11))
	c.Assert(err, IsNil)
	d2, err := s.E.Decrypt([]byte(e2))
	c.Assert(err, IsNil)
	d22, err := s.E.Decrypt([]byte(e2))
	c.Assert(err, IsNil)
	d3, err := s.E.Decrypt([]byte(e3))
	c.Assert(err, IsNil)

	c.Assert(s1, DeepEquals, string(d1))
	c.Assert(s1, DeepEquals, string(d11))
	c.Assert(s2, DeepEquals, string(d2))
	c.Assert(s2, DeepEquals, string(d22))
	c.Assert(s3, DeepEquals, string(d3))
}
