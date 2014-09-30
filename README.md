introduction-to-algorithms-reading
==================================

Summaries and exercises from my "Introduction to algorithms" reading

Chapters' explanations
======================

Introduction to start thinking how to design and analyze algorithms. It introduces how to specify algorithms, some design strategies and fundamental ideas in algorithm analysis.

Chapter 1 presents the algorithms in the context of modern computing systems. It provides an algorithm definition and consider some examples of algorithms as a technology.

The Role of Algorithms in Computing
===================================

Informal definition for algorithm
---------------------------------

An algorithm is a set of well-defined *computational steps* to apply on a set of *inputs* to produce a set of result values as *outputs*.

As an alternative definition we can view an algorithm as *tool* for solving a well-defined *computational problem*. The problem is defined in terms of the relation between a set of desired *outputs* derived from the provided *inputs*.

As a *instance of a problem* we know the *inputs* provided to fulfill the algorithm's constrains.

An algorithm is said to be correct if for every input *instance*, it halts with the correspondence *outputs*.

Data structures
---------------

> A data structure is a way to store and organize data in order to facilitate access and modifications. No single data structure works well for all purposes, and so it is important to know the strengths and limitations of several of them.

1.2 Algorithms as a technology
------------------------------

Efficiency: An efficient algorithm can make the difference regards to speed, when comparing two algorithms, even if the more efficient algorithm has a naive implementation.

In the example the raw computer power of A is 1,000 greater than B. But for a large array we can get a 17 times more speed with the B computer because we use `merge sort` algorithm that can handle more efficiently large arrays than `insert sort` on the A computer.

Algorithms and other technologies
---------------------------------

As we can see, the chosing of the algorithm as the hardware, affects the efficiency of ours programs, and this why we can think in the algorithms as a technologies.

> Having a solid base of algorithmic knowledge and technique is one characteristic that separates the truly skilled programmers from the novices. With modern com- puting technology, you can accomplish some tasks without knowing much about algorithms, but with a good background in algorithms, you can do much, much more.

Algorithms f(n) order by its speed
-----------------------------------

Faster:
* ln(n)
* sqrt(n)
* n
* n ln(2)
* n^2
* n^3
* 2^n
* n!

Analyzing algorithms
--------------------

> Analyzing an algorithm has come to mean predicting the resources that the algorithm requires. Occasionally, resources such as memory, communication bandwidth, or computer hardware are of primary concern, but most often it is computational time that we want to measure.
