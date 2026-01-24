import { memo } from "react";

import LeftContainer from "./LeftContainer";
import RightContainer from "./RightContainer";
import styles from "./styles.module.css";

const Header = memo(() => {
  return (
    <div className={styles.header}>
      {/* 左侧容器：Logo + 主题切换 */}
      <LeftContainer />

      {/* 右侧容器：操作按钮 */}
      <RightContainer />
    </div>
  );
});

Header.displayName = "Header";

export default Header;
