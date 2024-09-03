div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div mt-9 text-center text-2xl
        Existing Projects
      {{ range $i, $item := .items }} 
        div flex space-x-6
          div w-32 truncate
            {{ $item.name }}
          div
            {{ $item.file }}
      {{ end }}
      div mt-20 text-center text-2xl
        Start a New Project
      div mt-3 text-center
        form id=upload-form space-y-3 enctype=multipart/form-data
          div
            project name
          div
            input type=text id=name autofocus=true
          div
            video file (up to 500MB)
          div
            input type=file name=file id=file
          div mt-3
            input id=go type=submit value=go btn btn-primary
