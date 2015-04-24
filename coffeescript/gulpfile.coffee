# based on http://pem-musing.blogspot.com/2014/02/a-gulp-of-coffee-your-gulpfile-in.html
# and http://david.nowinsky.net/gulp-book/example/coffee.html

gulp = require 'gulp'
gutil = require 'gulp-util'
coffee = require 'gulp-coffee'
coffeelint = require 'gulp-coffeelint'

gulp.task 'lint', ->
    gulp.src './test.coffee'
        .pipe coffeelint()
        .pipe coffeelint.reporter()

gulp.task 'compile', ->
    gulp.src './test.coffee'
        .pipe coffee bare: true
        .pipe gulp.dest '.'
        .on 'error', gutil.log

gulp.task 'default', ['lint', 'compile']
