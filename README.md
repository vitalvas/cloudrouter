# CloudRouter

> **Warning**
> Current project under R&D. There is no guarantee that it will work, but corrections are welcome.

CloudRouter is a pure-GO implementation of small router.

Main idea - creating a router that does not use any third-party applications or any warpers. Only the linux kernel and applications from this project.

It is created both for clouds (on virtual machines) and for home use (EOM router).

## Services

* [X] DNS Proxy
  * [X] Static upstream
  * [ ] Cache
  * [ ] DNS over HTTPS
  * [ ] DNS over QUIC
* [X] Network Config
  * [ ] Interface config
  * [ ] Routing
  * [X] WireGuard VPN
* [X] Neighbor discovery
  * [X] Personal protocol (CRDP)
  * [ ] LLDP
