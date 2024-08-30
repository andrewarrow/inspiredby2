# inspiredby2

```
ffmpeg -i file1.mp4 -i file2.mp4 -filter_complex "[1:v]scale=1280:720[bg];[0:v]scale=320:180[pip];[bg][pip]overlay=W-w-10:10" -c:v libx264 -crf 23 -preset veryfast -movflags +faststart -y output.mp4


ffmpeg -i file1.mp4 -i file2.mp4 -filter_complex "[1:v]scale=1280:720[bg];[0:v]scale=320:180[pip];[bg][pip]overlay=W-w-10:10[v];[0:a][1:a]amix=inputs=2:duration=longest[a]" -map "[v]" -map "[a]" -c:v libx264 -crf 23 -preset veryfast -movflags +faststart -y output.mp4

ffmpeg -i file1.mp4 -i file2.mp4 -filter_complex "[1:v]scale=1280:720[bg];[0:v]scale=320:180,pad=360:220:20:20:orange[bordered_pip];[bg][bordered_pip]overlay=50:50[v];[0:a][1:a]amix=inputs=2:duration=longest[a]" -map "[v]" -map "[a]" -c:v libx264 -crf 23 -preset veryfast -movflags +faststart -y output.mp4



```

# silence hiss

```
ffmpeg -loop 1 -i frame005.jpg -f lavfi -i anoisesrc=d=10:c=pink -filter:a "volume=0.02" -c:v libx264 -t 10 -pix_fmt yuv420p -vf "scale=1280:720" -c:a aac -b:a 192k output.mp4
```


```
        <iframe width="560" height="315" src="https://www.youtube.com/embed/UfnAOcBirAs?si=V-DTImPbskGYJN9j" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
```

```
        <iframe width="560" height="315" src="https://www.youtube.com/embed/4voK6NiR4Xs?si=Z7tjdb3D3k-SNnFE" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
```
