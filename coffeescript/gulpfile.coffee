gulp = require 'gulp'
coffeelint = require 'gulp-coffeelint'

gulp.task 'lint', ->
    gulp.src './*.coffee'
        .pipe coffeelint()
        .pipe coffeelint.reporter()

gulp.task 'default', ['lint']
