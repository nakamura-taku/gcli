// DO NOT EDIT THIS FILE.
// THIS FILE IS GENERATED BY GO GENERATE.

/*
Command gcli generates a skeleon (codes and its directory structure) you need to start building CLI tool by Golang.
https://github.com/tcnksm/gcli

Usage:

    gcli [-version] [-help]  <command> [<options>]

Available commands:

    apply       Apply design template file for generating cli project
    design      Generate project design template
    list        List available cli frameworks
    new         Generate new cli project
    validate    Validate design template file

Use "gcli <command> -help" for more information about command.



Apply design template file for generating cli project

Apply design template file for generating cli project. You can generate
design template file via 'gcli design' command. If framework name is not
specified gcli use codegangsta/cli. You can set framework name via '-F'
option. To check cli framework you can use, run 'gcli list'. 

Usage:

  gcli apply [option] FILE


Options:

   -framework=name, -F        Cli framework name. By default, gcli use "codegangsta/cli"
                              To check cli framework you can use, run 'gcli list'.
                              If you set invalid framework, it will be failed.

   -skip-test, -T             Skip generating *_test.go file. By default, gcli generates
                              test file If you specify this flag, gcli will not generate
                              test files.



Generate project design template

Generate project design template (as toml file). You can pass that file to 'gcli apply'
command and generate CLI tool based on template file. You can define what command
and what flag you need on that file.

Usage:

  gcli design [option] NAME


Options:

  -command=name, -c           Command name which you want to add.
                              This is valid only when cli pacakge support commands.
                              This can be specified multiple times. Synopsis can be
                              set after ":". Namely, you can specify command by 
                              -command=NAME:SYNOPSYS. Only NAME is required.
                              You can set multiple variables at same time with ","
                              separator.

  -flag=name, -f              Global flag option name which you want to add.
                              This can be specified multiple times. By default, flag type
                              is string and its description is empty. You can set them,
                              with ":" separator. Namaly, you can specify flag by
                              -flag=NAME:TYPE:DESCIRPTION. Order must be flow  this and
                              TYPE must be string, bool or int. Only NAME is required.
                              You can set multiple variables at same time with ","
                              separator.

  -framework=name, -F         Cli framework name. By default, gcli use "codegangsta/cli"
                              To check cli framework you can use, run 'gcli list'.
                              If you set invalid framework, it will be failed.

  -owner=name, -o             Command owner (author) name. This value is also used for
                              import path name. By default, owner name is extracted from
                              ~/.gitconfig variable.


  -output, -O                 Change output file name. By default, gcli use "NAME-design.toml"



List available cli frameworks

Show all avairable cli frameworks.

Usage:

  gcli list



Generate new cli project

Generate new cli skeleton project. At least, you must provide executable
name. You can select cli package and set commands via command line option.
See more about that on Options section. By default, gcli use codegangsta/cli.
To check cli framework you can use, run 'gcli list'. 

Usage:

    gcli new [option] NAME

Options:

  -command=name, -c           Command name which you want to add.
                              This is valid only when cli pacakge support commands.
                              This can be specified multiple times. Synopsis can be
                              set after ":". Namely, you can specify command by 
                              -command=NAME:SYNOPSYS. Only NAME is required.
                              You can set multiple variables at same time with ","
                              separator.

  -flag=name, -f              Global flag option name which you want to add.
                              This can be specified multiple times. By default, flag type
                              is string and its description is empty. You can set them,
                              with ":" separator. Namaly, you can specify flag by
                              -flag=NAME:TYPE:DESCIRPTION. Order must be flow  this and
                              TYPE must be string, bool or int. Only NAME is required.
                              You can set multiple variables at same time with ","
                              separator.

   -framework=name, -F        Cli framework name. By default, gcli use "codegangsta/cli"
                              To check cli framework you can use, run 'gcli list'.
                              If you set invalid framework, it will be failed.

   -owner=name, -o            Command owner (author) name. This value is also used for
                              import path name. By default, owner name is extracted from
                              ~/.gitconfig variable.

   -skip-test, -T             Skip generating *_test.go file. By default, gcli generates
                              test file If you specify this flag, gcli will not generate
                              test files.

Examples:

To create todo command application skeleton which has 'add' and 'delete' command,

   $ gcli new -command=add:"Add new task" -command=delete:"delete task" todo



Validate design template file

Validate design template file which has required filed. If not it returns
error and non zero value. 

Usage:

  gcli validate FILE



*/
package main
