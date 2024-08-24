div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div text-white text-center text-2xl mt-3
        Welcome to InspiredBy2
      div mt-3
        You give us a youtube link and $6 USD and we turn it into the
        ultimate summary page with a link you can send to a busy person.
      div mt-3 bg-red-900 rounded-lg p-3
        There are lots of ways to pick which parts of the video are most important that should have a big part in the exec summary you are making. 
        Having to select the exact time codes in the video doesn't work very well.
        Instead we use AI to find some good candidates and then let you click
        next, next, next until you agree the AI found the perfect 1 min clip
that helps sell the whole summary.
      div mt-3 bg-purple-900 rounded-lg p-3
        But that's just the start, we do a lot more to, see demo. Or pony up
        six dollars and just go for it.
      div mt-3
        form id=welcome-form target=_blank action=https://buy.stripe.com/test_cN23e40qW2fA024cMM?prefilled_email!w method=GET
          div
            youtube link
          div
            input type=text id=link autofocus=true value=www.youtube.com/watch?v!Vcw1bYz3heI
          div
            your email
          div
            input type=text id=email value=oneone@gmail.com
          div mt-3
            input type=submit value=go btn btn-primary
