package constant

const(
	Host1 = "127.0.0.1:2379"
	Host2 = "127.0.0.1:22379"
	Host3 = "127.0.0.1:32379"
)

var (
	Cluster = []string{Host1, Host2, Host3}
)
