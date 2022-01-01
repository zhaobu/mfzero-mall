wsl -d Ubuntu -u root ip addr del $(ip addr show eth0 ^| grep 'inet\b' ^| awk '{print $2}' ^| head -n 1) dev eth0
wsl -d Ubuntu -u root ip addr add 10.133.236.150/24 broadcast 10.133.236.255 dev eth0
wsl -d Ubuntu -u root ip route add 0.0.0.0/0 via 10.133.236.1 dev eth0
wsl -d Ubuntu -u root echo nameserver 10.133.236.1 ^> /etc/resolv.conf
powershell -c "Get-NetAdapter 'vEthernet (WSL)' | Get-NetIPAddress | Remove-NetIPAddress -Confirm:$False; New-NetIPAddress -IPAddress 10.133.236.1 -PrefixLength 24 -InterfaceAlias 'vEthernet (WSL)'; Get-NetNat | ? Name -Eq WSLNat | Remove-NetNat -Confirm:$False; New-NetNat -Name WSLNat -InternalIPInterfaceAddressPrefix 10.133.236.0/24;"