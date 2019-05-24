## 题目

```
Requirements:

- Implement a driverless Car based on the provided interface that would be capable of moving forward into different directions and turning clockwise in a car park

- The car park would be rectangular in shape with configurable dimensions

- The Car should remember its position and orientation

- An exception will be thrown when the car moves outside the car park boundaries

- Evidence should be provided to demonstrate that the Car implementation meets the acceptance criteria

- Implement as a Maven module, if you need to use any library, use only those available in Maven Central.

 

Design considerations:

- The Car interface below is provided as a guideline and can be changed

- Future extensibility for new Car features should be kept in mind

- Consider OO modeling for the application

- Keep minimal and viable for the final workable application

 

Example:

- A simple illustration of the exercise with a car initially positioned at X = 1 and Y = 1 and facing North would look like this on a car park with dimension 4x4:

 

(Y)

   +---+---+---+---+

4  |   |   |   |   |

   +---+---+---+---+

3  |   |   |   |   |         N

   +---+---+---+---+     W <-|-> E

2  |   |   |   |   |         S

   +---+---+---+---+

1  | C |   |   |   |

   +---+---+---+---+

    1    2   3   4    (X)

 

Acceptance Criteria:

- Given the Car is in position X = 1 and Y = 1 and facing North, when the Car turns clockwise, then the Car is still in the same position but is now facing East

- Given the Car is in position X = 1 and Y = 1 and facing North, when the Car moves forward, then the Car is still facing North but is now in position X = 1 and Y = 2

- Given the Car is in position X = 1 and Y = 1 and facing East, when the Car moves forward, then the Car is still facing East but is now in position X = 2 and Y = 1

- Given the Car is in position X = 1 and Y = 1 and facing West, when the Car moves forward, then an exception is thrown

- Given the Car is in position X = 1 and Y = 1 and facing East, when the Car moves forward twice, then the Car is still facing East but is now in position X = 3 and Y = 1

 

 

Example Interface:

 

public interface Car {

   void move(String command);

    int getPositionX();

    int getPositionY();

    String getOrientation();

}

 

Notes:

·         Acceptance criteria is what we can verify the functionalities of the application. They should be easy, automatic, repeatable ran and built for maintainability purpose.
```


### 关于代码
写了一个简单和一个稍多一点抽象的两个版本。
两个版本里面都有一些让人看起来不爽的switch语句，先看着吧，有空再干掉它们。


---------------------------


### 关于测试
测试代码没写。


---------------------------


### 关于注释
没有怎么认真写，注释是代码的一部分，也相当重要。本人只是简单写了一下注释，代码不复杂，一看明了。


---------------------------


### 关于第二版的使用
第二版做成了交互式的，用户可以持续输入命令，命令一共有四个，分别是初始化停车场(p或park)、在当前方向上前进(f或forward)、顺时针转向90度(t或turn)和退出(q或quit),使用方法分别如下：


```
初始化停车场：p或park，初始化停车场，需要2个参数，分别是停车场的x值和y值

在当前方向上前进：f或forward，需要一个参数，指出向前移动多少

顺时针转向90度：t或turn，不需要参数

退出：q或quit，不需要参数
```

---------------------------

### 其他想法
本来想把channel、goroutine、web等东西也加进来玩的，特别是命令部分，可以设计模式里的命令模式来玩，还是先不玩了，时间有限，以后慢慢玩。
以后再把地图搞复杂些，加入障碍物啥的，哈哈哈。
