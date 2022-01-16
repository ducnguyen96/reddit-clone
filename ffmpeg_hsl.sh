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
    -map a:0 -c:a aac -b:a 128k -ac 1 -ar 44100\
    -f hls -hls_time $HLS_TIME -hls_playlist_type vod \
    -master_pl_name $1.m3u8 \
    -var_stream_map "v:0,a:0" $VIDEO_OUT/stream_%v.m3u8