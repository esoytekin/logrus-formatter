# logrus-formatter

console formatter for logrus

```
defaultLogFormat       = "%lvl% [ %time% ] %file% %func% => %msg%\n"
defaultTimestampFormat = "2006-01-02 15.04.05"
```

usage:

```
logrus.SetFormatter(&formatter.Formatter{
        TimestampFormat: "2006-01-02 15.04.05",
        LogFormat:       "%lvl% [ %time% ] %file% %func% => %msg%\n",
})

```
