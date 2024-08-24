div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div text-white text-center text-2xl mt-3
        Welcome to InspiredBy2
      div mt-3
        You give us a youtube link and $6 USD and we turn it into the
        ultimate summary page with a link you can send to a busy person.
      div mt-3
 
      div mt-3
        form
          div
            youtube link
          div
            input type=text id=link autofocus=true
          div mt-3
            input type=submit value=go btn btn-primary
