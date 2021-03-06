// Check for GO ODBC issue 23
// Ref: https://code.google.com/p/odbc/issues/detail?id=23

package main

import (
	_ "code.google.com/p/odbc"
	"database/sql"
	"fmt"
	"log"
)

func main() {
	conn, err := sql.Open("odbc",
		"driver=sql server;server=localhost;database=tempdb;trusted_connection=yes;")
	if err != nil {
		fmt.Println("Connecting Error")
		return
	}
	defer conn.Close()

	stmt, err := conn.Prepare("Select "+
		"'1' AS '(PDH-CSV 4.0) (Eastern Daylight Time)(240)', "+
		"'1' AS '\\\\MYDBSVR01\\Memory\\Available MBytes', "+
		"'1' AS '\\\\MYDBSVR01\\Memory\\Page Faults/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Memory\\Pages/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Memory\\% Committed Bytes In Use', "+
		"'1' AS '\\\\MYDBSVR01\\Memory\\Pool Paged Bytes', "+
		"'1' AS '\\\\MYDBSVR01\\Memory\\Pool Nonpaged Bytes', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Teredo Tunneling Pseudo-Interface)\\Current Bandwidth', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _38)\\Current Bandwidth', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _39)\\Current Bandwidth', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(isatap.dayforce.com)\\Current Bandwidth', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Teredo Tunneling Pseudo-Interface)\\Output Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _38)\\Output Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _39)\\Output Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(isatap.dayforce.com)\\Output Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Teredo Tunneling Pseudo-Interface)\\Bytes Total/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _38)\\Bytes Total/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _39)\\Bytes Total/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(isatap.dayforce.com)\\Bytes Total/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Teredo Tunneling Pseudo-Interface)\\Packets Sent/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _38)\\Packets Sent/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _39)\\Packets Sent/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(isatap.dayforce.com)\\Packets Sent/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Teredo Tunneling Pseudo-Interface)\\Bytes Sent/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _38)\\Bytes Sent/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _39)\\Bytes Sent/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(isatap.dayforce.com)\\Bytes Sent/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Teredo Tunneling Pseudo-Interface)\\Packets Received/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _38)\\Packets Received/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _39)\\Packets Received/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(isatap.dayforce.com)\\Packets Received/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Teredo Tunneling Pseudo-Interface)\\Bytes Received/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _38)\\Bytes Received/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(Broadcom BCM5708C NetXtreme II GigE [NDIS VBD Client] _39)\\Bytes Received/sec', "+
		"'1' AS '\\\\MYDBSVR01\\Network Interface(isatap.dayforce.com)\\Bytes Received/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Avg. Disk sec/Read', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Avg. Disk sec/Read', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Avg. Disk sec/Read', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Disk Writes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Disk Writes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Disk Writes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Disk Reads/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Disk Reads/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Disk Reads/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Disk Write Bytes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Disk Write Bytes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Disk Write Bytes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\% Disk Read Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\% Disk Read Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\% Disk Read Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Avg. Disk Bytes/Write', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Avg. Disk Bytes/Write', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Avg. Disk Bytes/Write', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Avg. Disk sec/Transfer', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Avg. Disk sec/Transfer', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Avg. Disk sec/Transfer', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Disk Transfers/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Disk Transfers/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Disk Transfers/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Disk Read Bytes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Disk Read Bytes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Disk Read Bytes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Disk Bytes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Disk Bytes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Disk Bytes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Avg. Disk Write Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Avg. Disk Write Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Avg. Disk Write Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\% Disk Write Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\% Disk Write Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\% Disk Write Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Current Disk Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Current Disk Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Current Disk Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Split IO/Sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Split IO/Sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Split IO/Sec', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\% Disk Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\% Disk Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\% Disk Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Avg. Disk Bytes/Read', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Avg. Disk Bytes/Read', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Avg. Disk Bytes/Read', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Avg. Disk Bytes/Transfer', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Avg. Disk Bytes/Transfer', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Avg. Disk Bytes/Transfer', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Avg. Disk Read Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Avg. Disk Read Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Avg. Disk Read Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Avg. Disk Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Avg. Disk Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Avg. Disk Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\% Idle Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\% Idle Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\% Idle Time', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(0 C: P: D:)\\Avg. Disk sec/Write', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(1 F:)\\Avg. Disk sec/Write', "+
		"'1' AS '\\\\MYDBSVR01\\PhysicalDisk(_Total)\\Avg. Disk sec/Write', "+
		"'1' AS '\\\\MYDBSVR01\\Processor(_Total)\\% Processor Time', "+
		"'1' AS '\\\\MYDBSVR01\\Processor(_Total)\\% Privileged Time', "+
		"'1' AS '\\\\MYDBSVR01\\Processor(_Total)\\% User Time', "+
		"'1' AS '\\\\MYDBSVR01\\System\\Processor Queue Length', "+
		"'1' AS '\\\\MYDBSVR01\\System\\Threads', "+
		"'1' AS '\\\\MYDBSVR01\\System\\Context Switches/sec', "+
		"'1' AS '\\\\MYDBSVR01\\System\\Processes', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlservr)\\Virtual Bytes', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlwriter)\\Virtual Bytes', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlservr)\\% User Time', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlwriter)\\% User Time', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlservr)\\Thread Count', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlwriter)\\Thread Count', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlservr)\\Working Set', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlwriter)\\Working Set', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlservr)\\% Privileged Time', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlwriter)\\% Privileged Time', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlservr)\\Private Bytes', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlwriter)\\Private Bytes', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlservr)\\Handle Count', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlwriter)\\Handle Count', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlservr)\\% Processor Time', "+
		"'1' AS '\\\\MYDBSVR01\\Process(sqlwriter)\\% Processor Time', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Access Methods\\Page Splits/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Access Methods\\Index Searches/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Access Methods\\Pages Allocated/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Access Methods\\Range Scans/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Access Methods\\Table Lock Escalations/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Access Methods\\Probe Scans/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Access Methods\\Full Scans/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(000)\\Free list empty/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(001)\\Free list empty/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(005)\\Free list empty/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(002)\\Free list empty/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(006)\\Free list empty/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(004)\\Free list empty/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(007)\\Free list empty/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(003)\\Free list empty/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(000)\\Free pages', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(001)\\Free pages', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(005)\\Free pages', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(002)\\Free pages', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(006)\\Free pages', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(004)\\Free pages', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(007)\\Free pages', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(003)\\Free pages', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(000)\\Free list requests/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(001)\\Free list requests/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(005)\\Free list requests/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(002)\\Free list requests/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(006)\\Free list requests/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(004)\\Free list requests/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(007)\\Free list requests/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Buffer Partition(003)\\Free list requests/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Databases(_Total)\\Log Flushes/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Databases(_Total)\\Transactions/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Databases(_Total)\\Log Cache Hit Ratio', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Databases(_Total)\\Active Transactions', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Databases(_Total)\\Log Cache Reads/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Databases(_Total)\\Log Truncations', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Databases(_Total)\\Log Flush Waits/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:General Statistics\\Logouts/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:General Statistics\\User Connections', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:General Statistics\\Logins/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Latches\\Average Latch Wait Time (ms)', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Latches\\Latch Waits/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Locks(_Total)\\Average Wait Time (ms)', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Locks(_Total)\\Lock Waits/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Locks(_Total)\\Lock Timeouts/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Locks(_Total)\\Number of Deadlocks/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:Locks(_Total)\\Lock Wait Time (ms)', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:SQL Statistics\\Failed Auto-Params/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:SQL Statistics\\SQL Re-Compilations/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:SQL Statistics\\SQL Compilations/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:SQL Statistics\\Unsafe Auto-Params/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:SQL Statistics\\Batch Requests/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:SQL Statistics\\Auto-Param Attempts/sec', "+
		"'1' AS '\\\\MYDBSVR01\\SQLServer:SQL Statistics\\Safe Auto-Params/sec'")
	if err != nil {
		log.Fatal(err)
	}

  row, err := stmt.Query()
	_ = row
  if err != nil {
    log.Fatal(err)
  }
}
