# Best practices

A set of best practices while building **Go** code.

### Project structure

Only separate pieces of application that can be used *on their own*. Keep related stuff in one package - it gets rid of import hell. If you need granularity in your package, split it up into more *files* instead.

Always import from **outside-in**:

```
app/
    main.go
    /module
        module.go
        /utils
        util.go
```

`main` imports `module`, that imports `utils`.

Write config functions for hooking packages up to each other *at run time* rather than *compile time*. Instead of `routes` importing all the packages that define routes, it can export `routes.Register`, which `main` (or code in each app) can call. [Source.](http://stackoverflow.com/questions/20380333)