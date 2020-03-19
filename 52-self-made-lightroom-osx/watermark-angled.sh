#!/usr/bin/env bash

# Input image: d.jpg, 960x1280px

# Since watermark goes diagonally, it's width *should* be:
# wm_width = sqrt(pow(width, 2) + pow(height, 2)) - some_delta;
# But, I noticed, it consumes CPU a lot when text goes out of the image area. Thus, I use average dimension:
# wm_width = (width + height) / 2 = (1280 + 960) / 2 = 1120

# Rotation angle should also depend on image width and height:
# wm_angle = arctg(width / height) - pi / 2 =  arctg(960 / 1280) - pi / 2
# 0.64350110879328 - 1.570796327 = −0.927295218 radians ~= -53 degrees

wm_width=1120
wm_angle=-53
wm_opacity=35
wm_text="WATERMARK_TEXT"
src_filename="d.jpg"
outfile="out.jpg"
wm_font="Menlo"

data=`convert -gravity center -fill grey75 -font $wm_font \
    -size $wm_width label:$wm_text \
    -format "%wx%hx%[label:pointsize]" info:`
echo "data:$data"
wd=`echo "$data" | cut -dx -f 1`
ht=`echo "$data" | cut -dx -f 2`
dim="${wd}x${ht}"
point=`echo "$data" | cut -dx -f 3`

convert $src_filename \( -size $dim xc:none \
    -gravity center -fill white -font $wm_font -pointsize $point -annotate +3+3 "$wm_text" \
    -gravity center -fill grey75 -font $wm_font -pointsize $point -annotate +0+0 "$wm_text" \
    -background none -rotate $wm_angle \) \
    -compose dissolve -define compose:args=$wm_opacity,100 -composite $outfile
