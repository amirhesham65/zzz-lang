# ZZZ-lang

ZZZ is an interpreted programming language developed in Go, featuring a syntax tailored for Gen Z. It supports variables, conditionals, first-class functions with closures, among other features.

I created this meme language as a fun project to deepen my understanding of programming language mechanics, inspired by the book [Writing An Interpreter In Go](https://interpreterbook.com/).

## Language Tour

### Variables

You can declare variables using the `lit` keyword, followed by the variable name and its value.

```zzz
lit cel = 42;
lit feh = cel * 9 / 5 + 32;
lit isCool = yea;
lit isLit = nah;
```

### Conditionals

You can use the `fr` keyword to start a conditional block, followed by the condition in parentheses. To define an alternate path or else block, use the `lowkey` keyword.

```zzz
lit age = 16;
fr (age > 18) {
    yea
} lowkey {
    nah
};
```

Note that conditionals or `fr` are expressions, meaning they return a value. In this case, the value of the block that gets executed. Also, conditionals in ZZZ has implicit returns.

```zzz
lit age = 16;
lit canDrink = fr (age > 18) { yea } lowkey { nah };
```

### Functions

You can define functions using the `fun` keyword, followed by the function name, parameters, and the function body.

Function calls are expressions, meaning they return a value. You can also pass functions as arguments to other functions (first-class functions).

```zzz
lit add = fun(a, b) {
    a + b;
};
add(1, 2);

lit applyTwice = fun(x, fn) {
    fn(fn(x));
};
applyTwice(2, fun(x) {x * x;});

fun(x){
    x * x;
}(2);
```

### Closures

Functions in ZZZ support closures, meaning they can access variables defined in their outer scope.

```zzz
lit add = fun(a) {
    fun(b) {
        a + b;
    };
};

add(2)(3);
```
