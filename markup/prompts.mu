div p-0 id=top
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2 space-y-6
      div mt-6 md:mt-0 font-allan text-4xl text-white uppercase
        Prompts
      {{ $items := .items }}
      {{ range $i, $item := $items }}
        form id=p-{{$item.guid}} flex space-x-3 prompt-form
          div w-32
            {{$item.section}}
          div w-full
            div
              {{$item.id_pika}}
            div
              <img src="{{$item.video_poster}}" class="w-96" />
            div
              {{$item.prompt_text}}
            div
              div flex space-x-6 items-center
                div id=d-{{$item.guid}}
                  {{$item.duration}}
                div
                  input type=hidden value={{$item.guid}} id=guid
                  input type=submit btn btn-sm btn-primary value=bump id=b-{{$item.guid}}
      {{ end }}
      
