# Diggler

Diggler is an open-source tool for gathering system information and its current status for forensic analysis. It is designed to be easier and faster when creating IBM PMR troubleshooting logs for IBM applications like WebSphere, DB2, MQ, MDM, BAW, IIS, etc. 

Diggler works on both Linux and Windows systems, making it a cross-platform tool. It gathers information such as:


## System & Environment

- OS details: Distro, kernel version, service pack level, patch history

- System uptime and boot time

- Time synchronization: NTP config, drift, clock skew (common in clusters)

- Environment variables (important for IBM tools, Java apps, DB clients)

- Java/JVM details: Installed JDKs, versions, JAVA_HOME, IBM JRE if present

- User accounts: Logged-in users, recent logins, sudoers, failed login attempts

- Open files / handles (Linux `lsof`, Windows handles count)


## Network & Connectivity

- Listening ports / bound services (netstat, ss)

- Firewall rules (iptables, nftables, Windows Firewall)

- Routing table

- DNS resolver config (/etc/resolv.conf, Windows DNS settings)

- Hosts file content (IBM often uses static mappings)

- TLS certificates (system trust store, app-specific keystores, expiry dates)


## Storage & Filesystem

- I/O performance stats (iostat, disk queue length, avg wait time)

- Swap usage (total, used, free, swappiness)

- Inode usage (common FS issue on Linux)

- Filesystem mounts with options (noexec, nodev, nosuid, etc.)


## Processes & Performance

- Top resource consumers (CPU, memory, I/O heavy processes)

- Process tree (to trace child processes like WebSphere JVMs, MQ daemons)

- Core dumps availability (ulimit -c, Windows WER config)

- Thread dumps (JVM apps) if enabled

- GC logs for WebSphere / Java workloads


## IBM Stack-Specific

- WebSphere: Profile configs, JVM logs, Deployment Manager sync status, DataSources, Thread pool config

- DB2: Instance config, bufferpool size, db2diag logs, db2pd snapshots

- MQ: Queue manager status, channel status, qmgr logs, error logs

- MDM/BAW/IIS: Product version, fix pack level, specific logs (e.g. MDM logs in `/opt/IBM/MDM/logs`)

- Installed Fixpacks / iFixes (pull from IBM product inventory commands)


## System Health Indicators

- Error logs summary (last N critical entries from syslog/event log)

- OOM events (Linux `dmesg`, Windows EventID 2004)

- Kernel panics / bugchecks

- Service status (systemd / Windows services, esp. IBM-related daemons)


## Security & Compliance

- SELinux/AppArmor status (Linux)

- Antivirus/Endpoint agents (may interfere with IBM apps)

- File permissions anomalies (common IBM PMR troubleshooting)

- SSL/TLS configs for IBM apps (default ciphers, keystores)


## Optional (Deep Forensic)

- Hash of binaries/configs (to detect tampering)

- Package manager state (rpm -Va / debsums for integrity check)

- Crash dump configs (kdump, Windows crash dumps)
