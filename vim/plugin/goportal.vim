if !exists('g:goportal_default_fifoname')
    let g:goportal_default_fifoname = ""
endif

function! GoPortal(...)
    let cmd = join(a:000," ")
    call system("goportal ".cmd)
endfunction

function! GoPortalDefault(name)
    let g:goportal_default_fifoname = a:name
endfunction

command! -nargs=+ GoPortal   call GoPortal(g:goportal_default_fifoname,<q-args>)
command! -nargs=+ GoPortalInternal  call GoPortal("-i",g:goportal_default_fifoname,<q-args>)
