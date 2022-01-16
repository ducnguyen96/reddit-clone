#!/bin/sh
VIDEO_IN=/media/videos/raw/$1
VIDEO_OUT=/media/videos/dash/$1
HLS_TIME=4
FPS=25
GOP_SIZE=100
CRF_P=21
PRESET_P=veryfast
V_SIZE_0=640x360

mkdir /media/videos/dash/$1

ffmpeg -i $VIDEO_IN \
    -preset $PRESET_P -keyint_min $GOP_SIZE -g $GOP_SIZE -sc_threshold 0 \
    -r $FPS -c:v libx264 -pix_fmt yuv420p -c:a aac -b:a 128k -ac 1 -ar 44100 \
    -map v:0 -s:0 $V_SIZE_0 -maxrate:0 365k -bufsize:0 730k \
    -map 0:a \
    -init_seg_name init\$RepresentationID\$.\$ext\$ -media_seg_name chunk\$RepresentationID\$-\$Number%05d\$.\$ext\$ \
    -use_template 1 -use_timeline 1  \
    -seg_duration 4 -adaptation_sets "id=0,streams=v id=1,streams=a" \
    -f dash $VIDEO_OUT/$1.mpd