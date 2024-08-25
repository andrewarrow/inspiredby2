div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div text-white text-center text-2xl mt-3
        Success!
      div mt-3 bg-purple-900 rounded-lg p-3
        Thank you for your payment.
      div mt-3 bg-orange-900 rounded-lg p-3 text-white
        div
          We are now downloading your link: {{.link}}
        div
          Stay on this page...
      div hidden
        div id=guid
          {{.guid}}
      div mt-3 id=duration
      div mt-3 bg-grey-900 rounded-lg p-3 text-white
        div 
          <img src="" id="photo1" class="w-64 rounded-lg"/>
        div 
          <img src="" id="photo2" class="w-64 rounded-lg"/>
      div id=canvas
   
