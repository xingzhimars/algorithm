## Trie树

rie树（前缀树、字典树）用于快速查询某个字符串或者字符前缀是否存在的数据结构。

作用：高效地**存储**和**查找**字符串集合的数据结构



凡是用Trie树，一般题目中肯定限制了字母只有26或者52个。

一般存的话，会在结尾的单词打上一个结束的标记



## 节点方案

```java
class Trie {

    class TrieNode {
        boolean end;
        TrieNode[] children = new TrieNode[26];
    }

    TrieNode root;

    Trie() {
        root = new TrieNode();
    }

    // p指针是遍历用的
    void insert(String word) {
        TrieNode p = root;
        for (int i = 0; i < word.length(); i++) {
            int index = word.charAt(i) - 'a';
            TrieNode child = p.children[index];
            if (child == null) {
                child = new TrieNode();
                p.children[index] = child;
            }
            p = child;
        }
        p.end = true;
    }

    boolean search(String word) {
        TrieNode p = root;
        for (int i = 0; i < word.length(); i++) {
            int index = word.charAt(i) - 'a';
            TrieNode child = p.children[index];
            if (child == null) return false;
            p = child;
        }
        return p.end;
    }

    boolean startsWith(String prefix) {
        TrieNode p = root;
        for (int i = 0; i < prefix.length(); i++) {
            int index = prefix.charAt(i) - 'a';
            TrieNode child = p.children[index];
            if (child == null) return false;
            p = child;
        }
        return true;
    }

}
```



## 数组方案

```go
package main

const N = 30010 * 3

var son [][]int // 存储的是每个节点的所有儿子
var cnt []int   // 以当前这个节点结尾的单词有多少个
var idx int     // 当前用到了哪个下标，下标为0的点，既是根节点，又是空节点

type Trie struct {
}

func Constructor() Trie {
        son = make([][]int, N)
        for i := 0; i < N; i++ {
                son[i] = make([]int, 26)
        }

        cnt = make([]int, N)

        return Trie{}
}

func (this *Trie) Insert(word string) {
        // 下标为0的点是根节点
        p := 0
        for i := 0; i < len(word); i++ {
                u := word[i] - 'a'
                if son[p][u] == 0 { // 不存在节点u
                        idx++
                        son[p][u] = idx // 保存节点u
                }
                p = son[p][u]
        }
        // 以p结尾的节点数
        cnt[p]++
}

func (this *Trie) Search(word string) bool {
        p := 0
        for i := 0; i < len(word); i++ {
                u := word[i] - 'a'
                if son[p][u] == 0 {
                        return false
                }
                p = son[p][u]
        }
        return cnt[p] > 0
}

func (this *Trie) StartsWith(prefix string) bool {
        p := 0
        for i := 0; i < len(prefix); i++ {
                u := prefix[i] - 'a'
                if son[p][u] == 0 {
                        return false
                }
                p = son[p][u]
        }
        return true
}
```

