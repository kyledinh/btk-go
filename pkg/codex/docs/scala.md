# Scala Docs 

## Generate a new Scala 3 Project 
- ` sbt new scala/scala3.g8` 

Generates:
```
$ tree scala-3-project-template 
scala-3-project-template
├── build.sbt
├── project
│   └── build.properties
├── README.md
└── src
    ├── main
    │   └── scala
    │       └── Main.scala
    └── test
        └── scala
            └── Test1.scala

```

## Sample `sbt.build` for testing 
```
name := "HelloScalaTest"
version := "0.1"
scalaVersion := "3.2.0"

libraryDependencies ++= Seq(
  "org.scalatest" %% "scalatest" % "3.2.9" % Test
)
```
