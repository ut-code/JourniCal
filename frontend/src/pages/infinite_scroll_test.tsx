import React, { useCallback, useRef, useState } from "react";

const InfiniteScrollTest: React.FC = () => {
  //表示するデータ
  const [list, setList] = useState<number[]>([]);
  const [bottomIndex, setBottomIndex] = useState<number>(0);
  const [topIndex, setTopIndex] = useState<number>(0);
  const [InitialScrollHeight, setInitialScrollHeight] = useState<number>(0);

  const listRef = useRef<HTMLDivElement>(null!);

  //項目を読み込むときのコールバック
  const topRef = useCallback(
    (element: HTMLDivElement | null) => {
      if (element === null) return;
      if (!listRef.current) return;
      setInitialScrollHeight(listRef.current.scrollHeight);
      const options = { threshold: 0.01 };
      const observer = new IntersectionObserver((entries, observer) => {
        const topRatio = entries[0].isIntersecting
        if (topRatio) {
          observer.disconnect();
          // topRef の DOM 要素が少しでも見えたらこの部分が実行
          setList([topIndex, ...list]);
          setTopIndex(topIndex - 1);
          // 要素追加前のスクロール位置に戻る
          window.scrollTo({
            top: listRef.current.scrollHeight - InitialScrollHeight,
          });
        }
      }, options);
      observer.observe(element);
    },
    [InitialScrollHeight, list, topIndex],
  );

  const bottomRef = useCallback(
    (element: HTMLDivElement | null) => {
      if (element === null) return;
      const options = { threshold: 0.01 };
      const observer = new IntersectionObserver((entries, observer) => {
        const bottomRatio = entries[0].intersectionRatio;
        if (bottomRatio > 0 && bottomRatio <= 1) {
          observer.disconnect();
          // bottomRef の DOM 要素が少しでも見えたらこの部分が実行
          setList([...list, bottomIndex]);
          setBottomIndex(bottomIndex + 1);
        }
      }, options);
      observer.observe(element);
    },
    [bottomIndex, list],
  );

  return (
    <div>
      <div ref={topRef} />
      <div ref={listRef}>
        {list.map((value) => (
          <div key={Math.random()}>{value}</div>
        ))}
      </div>
      <div ref={bottomRef} />
    </div>
  );
};

export default InfiniteScrollTest;
