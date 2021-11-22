def StringChallenge(timePoints):
    times = []
    l = len(timePoints)
    for t in timePoints:
        hours, min_with_day_part = t.split(":")

        hours = int(hours)
        if min_with_day_part[2:] == "pm":
            hours += 12
        mins = int(min_with_day_part[0:2])
        times.append(hours * 60 + mins)
    times.sort()

    res = min(times[-1] - times[0], 1440 - times[-1] + times[0])
    for i in range(0, l - 1):
        res = min(res, min(times[i + 1] - times[i], 1440 - times[i + 1] + times[i]))

    return str(res)


# keep this function call here
print(StringChallenge(["1:10pm", "4:40am", "5:00pm"]))
