div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div text-white text-center text-2xl mt-3
        Welcome to InspiredBy
      div mt-3
      div mt-9 text-center
        a href=/core/register btn btn-primary
          Get Started
