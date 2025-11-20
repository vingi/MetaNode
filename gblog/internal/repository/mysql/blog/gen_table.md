#### gblog.blog 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | AID |  | int(11) | PRI | NO | auto_increment |  |
| 2 | Author |  | varchar(49) |  | YES |  | NULL |
| 3 | Category |  | varchar(100) |  | YES |  | NULL |
| 4 | Title |  | varchar(100) |  | YES |  | NULL |
| 5 | Content |  | longtext |  | YES |  | NULL |
| 6 | Status |  | tinyint(4) |  | YES |  | NULL |
| 7 | AddTime | getdate() | datetime |  | YES |  | NULL |
| 8 | UpdateTime | getdate() | datetime |  | YES |  | NULL |
