div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      a href=/frame/routing 
        <img src="https://i.imgur.com/Dza3fNU.png" class="rounded-lg border-2" style="transform: rotate(-9deg);"/>
    div mt-6 md:mt-0 w-full md:w-1/2 font-allan text-4xl text-white uppercase
      div w-3/4 mt-9
        a href=/frame/routing 
          div text-center bg-white hover:bg-red-600 text-black rounded-lg p-3
            Heart Rate Variability (HRV)
      div w-3/4 mt-9 ml-9
        a href=/frame/sql
          div text-center bg-indigo-200 hover:bg-red-600 text-black rounded-lg p-3
            Oscillatory variability in milliseconds
      div w-3/4 mt-9 ml-18
        a href=/frame/migrations
          div text-center bg-rose-200 hover:bg-red-600 text-black rounded-lg p-3
            It's a periodic pattern
      div w-3/4 mt-9 ml-9
        a href=/frame/wasm 
          div text-center bg-yellow-200 hover:bg-red-600 text-black rounded-lg p-3
            Effectively respond to stressors
  div w-full font-familjen text-4xl mt-16 ml-0 md:ml-9
    div flex flex-col md:flex-row space-x-6
      div
        img w-64 rounded-md src=https://i.imgur.com/rmvpjcP.png
      div
        div text-white
          by Chris Williamson & Dr. Leah Lagos
        div text-2xl
          How To Improve Your Heart Rate Variability
        div flex justify-center py-9 p-3 md:space-x-6 
          a href=https://github.com/andrewarrow/
            div bg-white rounded-full p-3 w-24 h-24 flex items-center justify-center
              <svg xmlns="http://www.w3.org/2000/svg" width="60" height="60" viewBox="0 0 24 24" id="github"><path d="M12 .14c-6.617 0-12 5.383-12 12 0 5.576 3.95 10.5 9.392 11.708A.5.5 0 0 0 10 23.36v-2.72a.5.5 0 0 0-.5-.5h-1c-1.248 0-2.097-1.183-2.847-2.226-.077-.107-.153-.214-.23-.318.355.188.66.415.961.638.599.444 1.219.903 2.043.903h.01c.064.007.648.067 1.1-.33.212-.187.463-.54.463-1.167v-.349a.499.499 0 0 0-.364-.481C6.863 16.025 5 13.947 5 11.64c0-1.2.493-2.345 1.425-3.312a.5.5 0 0 0 .094-.558c-.372-.802-.293-1.893.148-2.548.584.227 1.341.704 1.833 1.288a.5.5 0 0 0 .554.147 8.67 8.67 0 0 1 5.893 0 .5.5 0 0 0 .554-.147c.492-.584 1.249-1.061 1.832-1.289.442.655.521 1.747.148 2.549a.5.5 0 0 0 .094.558C18.507 9.295 19 10.44 19 11.64c0 2.422-2.07 4.591-5.033 5.274a.5.5 0 0 0-.329.72c.247.47.362 1.107.362 2.006v3.72a.5.5 0 0 0 .608.488C20.05 22.64 24 17.716 24 12.14c0-6.617-5.383-12-12-12z"></path></svg>
          a href=https://www.instagram.com/andrewarrow/
            div bg-white rounded-full p-3 w-24 h-24 flex items-center justify-center
              <svg xmlns="http://www.w3.org/2000/svg" width="60" height="60" viewBox="0 0 16 16" id="instagram"><path d="M11 0H5a5 5 0 0 0-5 5v6a5 5 0 0 0 5 5h6a5 5 0 0 0 5-5V5a5 5 0 0 0-5-5zm3.5 11c0 1.93-1.57 3.5-3.5 3.5H5c-1.93 0-3.5-1.57-3.5-3.5V5c0-1.93 1.57-3.5 3.5-3.5h6c1.93 0 3.5 1.57 3.5 3.5v6z"></path><path d="M8 4a4 4 0 1 0 0 8 4 4 0 0 0 0-8zm0 6.5A2.503 2.503 0 0 1 5.5 8c0-1.379 1.122-2.5 2.5-2.5s2.5 1.121 2.5 2.5c0 1.378-1.122 2.5-2.5 2.5z"></path><circle cx="12.3" cy="3.7" r=".533"></circle></svg>
          a href=https://x.com/andrewarrow
            div bg-white rounded-full p-3 w-24 h-24 flex items-center justify-center
              <svg xmlns="http://www.w3.org/2000/svg" width="50" height="50" viewBox="0 0 1668.56 1221.19" id="twitter-x"><path d="M283.94,167.31l386.39,516.64L281.5,1104h87.51l340.42-367.76L984.48,1104h297.8L874.15,558.3l361.92-390.99 h-87.51l-313.51,338.7l-253.31-338.7H283.94z M412.63,231.77h136.81l604.13,807.76h-136.81L412.63,231.77z" transform="translate(52.39 -25.059)"></path></svg>
          a href=https://www.linkedin.com/in/andrewarrow
            div bg-white rounded-full p-3 w-24 h-24 flex items-center justify-center
              <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 5 1036 990" id="linkedin"><path d="M0 120c0-33.334 11.667-60.834 35-82.5C58.333 15.833 88.667 5 126 5c36.667 0 66.333 10.666 89 32 23.333 22 35 50.666 35 86 0 32-11.333 58.666-34 80-23.333 22-54 33-92 33h-1c-36.667 0-66.333-11-89-33S0 153.333 0 120zm13 875V327h222v668H13zm345 0h222V622c0-23.334 2.667-41.334 8-54 9.333-22.667 23.5-41.834 42.5-57.5 19-15.667 42.833-23.5 71.5-23.5 74.667 0 112 50.333 112 151v357h222V612c0-98.667-23.333-173.5-70-224.5S857.667 311 781 311c-86 0-153 37-201 111v2h-1l1-2v-95H358c1.333 21.333 2 87.666 2 199 0 111.333-.667 267.666-2 469z"></path></svg>
        div text-base font-mono
          a hover:underline href=https://youtu.be/wh_M25S2xUw
            [Original Hour Long Uncut Youtube](https://youtu.be/wh_M25S2xUw)
        div text-base font-mono mt-3
          a hover:underline href=https://github.com/andrewarrow/inspiredby2
            [Made With](https://github.com/andrewarrow/inspiredby2)
        div mt-3
          a href=/
            div text-center bg-green-600 text-black hover:bg-green-300 font-allan text-4xl uppercase rounded-lg p-3
              Experience Exec Summary
  div mt-9 w-full font-allan text-4xl text-white uppercase
    a href=/frame/templates
      img bg-cover bg-no-repeat w-full border-2 src=https://i.imgur.com/nTwdcaD.png
  div mt-16 mb-16 flex space-x-9 items-start justify-center mx-16
    div w-1/2 
      a href=/frame/migrations
        div font-allan text-4xl text-white uppercase 
          The nervous system, under stress, fragments autotomic bandwidth
    div w-1/2
      a href=/frame/migrations
        <img src="https://i.imgur.com/h55A0wQ.png" class="rounded-lg border-2" style="transform: rotate(9deg);"/>
  div mt-16 mb-16 flex space-x-9 items-start justify-center mx-16
    <iframe width="560" height="315" src="https://www.youtube.com/embed/KU6-BTxQoCA?si=oX8HhdG6lSKJYoRu" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
