# Itinerary Prettifier
a command line tool, which reads a text-based itinerary from a file (input), processes the text to make it customer-friendly, and writes the result to a new file (output).



## Usage 

The tool will be launched from the command line with three arguments:

1. Path to the input
2. Path to the output
3. Path to the airport lookup

```shell
$ go run . ./input.txt ./output.txt ./airports_lookup.csv
```



## Example

**input.txt:**

```

D(2022-05-09T08:07Z)


T12(2022-05-09T08:07Z)




T24(2022-05-09T08:07Z)








Your flight departs from #HAJ       at D(2022-05-09T08:07Z) ,            T12(2022-05-09T08:07Z),          and your destination is ##EDDW.



##AGGH



Hope\ryou\vhave\fa\n nice trip!



Some examples of time-date prettifier:




1. D(2022-05-09T08:07Z)
2. T12(2069-04-24T19:18-02:00)
3. T12(2080-05-04T14:54Z)
4. T12(1980-02-17T03:30+11:00)
5. T12(2029-09-04T03:09Z)
6. T24(2032-07-17T04:08+13:00)
7. T24(2084-04-13T17:54Z)
8. T24(2024-07-23T15:29-11:00)
9. T24(2042-09-01T21:43Z)

```

**output.txt:**

```

09 May 2022

08:07AM (+00:00)

08:07 (+00:00)

Your flight departs from Hannover Airport at 09 May 2022 , 08:07AM (+00:00), and your destination is Bremen Airport.

Honiara International Airport

Hope
you
have
a
nice trip!

Some examples of time-date prettifier:

1. 09 May 2022
2. 07:18PM (-02:00)
3. 02:54PM (+00:00)
4. 03:30AM (+11:00)
5. 03:09AM (+00:00)
6. 04:08 (+13:00)
7. 17:54 (+00:00)
8. 15:29 (-11:00)
9. 21:43 (+00:00)

```





