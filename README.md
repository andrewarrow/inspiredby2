# inspiredby2

```
ffmpeg -i file1.mp4 -i file2.mp4 -filter_complex "[1:v]scale=1280:720[bg];[0:v]scale=320:180[pip];[bg][pip]overlay=W-w-10:10" -c:v libx264 -crf 23 -preset veryfast -movflags +faststart -y output.mp4

```


```
        <iframe width="560" height="315" src="https://www.youtube.com/embed/UfnAOcBirAs?si=V-DTImPbskGYJN9j" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
```

```
        <iframe width="560" height="315" src="https://www.youtube.com/embed/4voK6NiR4Xs?si=Z7tjdb3D3k-SNnFE" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
```
