if !exists('g:goportal_default_fifoname')
    let g:goportal_default_fifoname = ""
endif

function! GoPortal(...)
    let cmd = g:goportal_default_fifoname." ".join(a:000," ")
    call system("goportal ".cmd)
endfunction

function! GoPortalDefault(name)
    let g:goportal_default_fifoname = a:name
endfunction

command! -nargs=+ GoPortal   call GoPortal(<q-args>)
command! -nargs=+ GoPortalInternal  call GoPortal("-i",<q-args>)
