div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div text-white text-center text-2xl mt-3
        Welcome to InspiredBy2
      div mt-3
        You give us a youtube link and $6 USD and we turn it into the
        ultimate summary video.
      div mt-3 bg-red-900 rounded-lg p-3
        We download the entire video. (Max length is 2 hours.)
        We break it into N segments one for each minute.
        Within each minute we break that into six, ten second files.
        We create summaries for each minute, and then summaries for the
        entire first third, middle third, and final third. i.e. the begining, 
        middle and end.
      div mt-3 bg-purple-900 rounded-lg p-3
        span
          To see what the process looks like, see
        span
          a href=/core/demo link
            our demo.
      div mt-3
        form id=welcome-form
          div
            youtube link
          div
            <input type="text" id="link" autofocus="true" value="https://youtu.be/wh_M25S2xUw"/>
          div
            your email
          div
            input type=text id=email value=oneone@gmail.com
          div mt-3
            input id=go type=submit value=go btn btn-primary
