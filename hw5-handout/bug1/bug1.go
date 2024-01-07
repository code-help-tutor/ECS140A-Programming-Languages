WeChat: cstutorcs
QQ: 749389476
Email: tutorcs@163.com
package bug1

// Counter stores a count.
type Counter struct {
	n int64
}

// Inc increments the count in the Counter.
func (c *Counter) Inc() {
	c.n++
}
