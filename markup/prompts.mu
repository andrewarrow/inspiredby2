div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2 space-y-6
      div mt-6 md:mt-0 font-allan text-4xl text-white uppercase
        Prompts
      {{ $items := .items }}
      {{ range $i, $item := $items }}
        div
          Hi {{$item}}
      {{ end }}
      
