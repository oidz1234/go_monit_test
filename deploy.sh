#!/bin/bash


# now set in env
api_key=$api_key
echo "api key is"
echo $api_key
deploy_location="/opt/monitor-monkey"
agent_location="https://github.com/oidz1234/go_monit_test/raw/master/monitor-monkey-agent"
agent="/opt/monitor-monkey/monitor-monkey-agent"
unit_file="/etc/systemd/system/monitor-monkey.service"
unit_name="monitor-monkey.service"

mkdir "$deploy_location"
cd "$deploy_location"
wget "$agent_location"
chmod +x "$agent"


unit_content="[Unit]
Description=Monitor Monkey Agent
After=network.target

[Service]
ExecStart=/opt/monitor-monkey/monitor-monkey-agent $api_key
Restart=always
RestartSec=5s

[Install]
WantedBy=default.target
"

echo "$unit_content" | sudo tee "$unit_file" > /dev/null
systemctl daemon-reload
systemctl start "$unit_name"

