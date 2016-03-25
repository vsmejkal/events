#!/usr/bin/env bash

apt-get update

# PostgreSQL
apt-get install postgresql postgresql-contrib
sudo -u postgres /bin/bash - << EOF
createuser www
createdb -O www eventsdb
EOF

# Automatic security updates
apt-get install unattended-upgrades
cat > /etc/apt/apt.conf.d/10periodic << EOF
APT::Periodic::Update-Package-Lists "1";
APT::Periodic::Download-Upgradeable-Packages "1";
APT::Periodic::AutocleanInterval "7";
APT::Periodic::Unattended-Upgrade "1";
EOF

