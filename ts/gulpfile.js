var gulp = require('gulp'),
    ts = require('gulp-typescript'),
    minifycss = require('gulp-minify-css'),
    concat = require('gulp-concat'),
    rename = require('gulp-rename'),
    uglify = require('gulp-uglify'),
    del = require('del');

var tsProject = ts.createProject('tsconfig.json');

gulp.task('combine', ['minifyjs'], function() {
    gulp.src('build/js/*.js').pipe(concat('main.js')).pipe(gulp.dest('build/js'));
    gulp.src('dist/js/*.js').pipe(concat('main.min.js')).pipe(gulp.dest('dist/js'));
})
  

gulp.task('minifyjs', ['build'], function() {
    return gulp.src('build/js/*.js')
        .pipe(rename({suffix:'.min'}))
        .pipe(uglify())
        .pipe(gulp.dest('dist/js'));  //输出
});


gulp.task('build', function() {
     return gulp.src('src/ts/*.ts')
        .pipe(tsProject())
        .pipe(gulp.dest('build/js'));          
});


gulp.task('clean', function(){
    del(['dist/js/*.js'],{force: true});
    del(['build/js/*.js'],{force: true});
 });        //清除dist中所有的文件和文件夹，适合新的项目开始时使用


gulp.task('default', ['clean'], function() {
    gulp.start('build','minifyjs','combine');
})