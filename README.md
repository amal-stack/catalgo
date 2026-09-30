# Catalgo

A CLI-driven tool to organize problem solving history into a structured knowledge base.

Catalgo (catalog + algorithm) is a data-driven CLI and knowledge base for organizing algorithmic problems, solutions, implementations, concepts, and their relationships across programming languages and platforms.


## Why I'm building Catalgo
As a software engineer, solving problems is the primary foundation of my work. It is often said that there is a disconnect between grinding competitive programming problems and solving real-world problems. I believe this is amplified by how algorithmic problem solving platforms are designed: the priorities are usually leaderboards, gamification, streaks, time-limits and expected outputs. This helps stay motivated but real-world demands differ. What helps there are pattern matching, recognizing trade-offs and constraints, and knowing when an algorithm or data structure is relevant.

Moreover, personally what was the most frustrating for me was a missing paper-trail. Over the years, I have spent time solving problems on several platforms but when I decided to look back, I had lost track of the problems I had solved and the things I had learnt: my learning was fragmented, and I had no idea how far I had come and even how many sites I had created an account on.

Catalgo aims to bridge this gap by treating each problem as part of a larger whole. Instead of seeing a problem as an isolated challenge in a checklist, Catalgo helps you document the associated concepts, patterns, relationships to other problems and your personal insights alongside the problem. It also aims to decouple a problem from its solutions and implementation styles. 

## Design and Concepts
Catalgo is an early-stage project. The ideas and data model are still being explored, and the interface is expected to evolve.

At its core, Catalgo would be a CLI tool that exposes commands to scaffold the structure for problems, associate it to patterns, and add notes to it. I currently plan to use YAML to model a problem. For example:
```yaml
id: 1
title: Contains Duplicate
difficulty: easy
slug: contains-duplicate

platforms:
  leetcode:
    id: 217 
  neetcode:
    slug: duplicate-integer

collections:
  - name: neetcode150

classification:
  patterns:
    - duplicate-detection
  data_structures:
    - array
    - hash-set
  algorithms:
    - sorting
  categories:
  tags:


solutions:
  - title: Brute force (compare every i from 0 to n to i+1 to n)
    complexity: 
      time: O(n^2)
      space: O(1)

  - title: Sort and compare adjacent
    complexity: 
      time: O(n log n)
      space: O(1)
  
  - title: Hash set
    complexity:
      time: O(n)
      space: O(n) 
```


