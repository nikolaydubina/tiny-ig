intall Go https://go.dev/doc/install

install this tool
```bash
go install github.com/nikolaydubina/tiny-ig@latest
```

go to directory with photos and run
```bash
tiny-ig -prefix="<where your photos stored here>" -w 304 -h 304 > gallery.md
```

example
```bash
tiny-ig -prefix="https://d3g5k8a9nzlkmy.cloudfront.net/photos/" -w 304 -h 304 > gallery.md
```

```bash
tiny-ig -prefix="https://nikolaydubina-blog-public.s3.ap-southeast-1.amazonaws.com/photos/" -w 304 -h 304 > gallery.md
```
