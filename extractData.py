import pandas as pd
import matplotlib.pyplot as plt

# df = pd.read_csv("datas/inputsize.csv")
# input_size = df["inputsize"]
# rounds = df["rounds"]
# time = df["time"]


# plt.plot(input_size, time, color = "red", label = "time")
# plt.plot(input_size, rounds, color = "blue", label = "rounds")
# plt.title('rounds,time vs input_size')
# plt.legend()
# plt.show()

# df = pd.read_csv("datas/fNode.csv")
# fNum = df["fNode"]
# rounds = df["rounds"]
# time = df["time"]


# plt.plot(fNum, time, color = "red", label = "time")
# plt.plot(fNum, rounds, color = "blue", label = "rounds")
# plt.title('rounds,time vs failureNode')
# plt.legend()
# plt.show()


# df = pd.read_csv("datas/minDelay.csv")
# mindelay = df["minDelay"]
# rounds = df["rounds"]
# time = df["time"]


# plt.plot(mindelay, time, color = "red", label = "time")
# plt.plot(mindelay, rounds, color = "blue", label = "rounds")
# plt.title('rounds,time vs mindelay')
# plt.legend()
# plt.show()

df = pd.read_csv("datas/maxDelay.csv")
maxdelay = df["maxDelay"]
rounds = df["rounds"]
time = df["time"]


plt.plot(maxdelay, time, color = "red", label = "time")
plt.plot(maxdelay, rounds, color = "blue", label = "rounds")
plt.title('rounds,time vs maxdelay')
plt.legend()
plt.show()