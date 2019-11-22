#! /bin/bash
# Build web UI

cd ~/my_go/video_server/web
go install
cp ~/go/bin/web ~/go/bin/video_server_web_ui
cp -R ~/my_go/video_server/templates ~/go/bin/video_server_web_ui/