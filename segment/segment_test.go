package segment

import (
	"github.com/forcemeter/chinese-segment-go/dicts"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
	"testing"
)

var text1 = `75年肝胆相照，75年荣辱与共。1949年9月21日，在中国人民争取民族独立和人民解放取得历史性胜利的凯歌声中，中国人民政治协商会议第一届全体会议开幕。人民政协是中国共产党把马克思列宁主义统一战线理论、政党理论、民主政治理论同中国实际相结合的伟大成果，是中国共产党领导各民主党派、无党派人士、人民团体和各族各界人士在政治制度上进行的伟大创造。75年来，在中国共产党的领导下，人民政协积极投身建立新中国、建设新中国、探索改革路的伟大实践，走过了辉煌历程。进入新时代，人民政协坚持以习近平新时代中国特色社会主义思想为指导，学深悟透习近平总书记关于加强和改进人民政协工作的重要思想，坚持稳中求进工作总基调，坚持党的领导、统一战线、协商民主有机结合，坚持团结和民主两大主题，发扬优良传统，牢记政治责任，提高政治协商、民主监督、参政议政水平，始终在党和国家工作大局下谋划和开展工作，紧紧围绕推进中国式现代化履行职能、凝心聚力，在实现中国梦的伟大实践中，创造了新的辉煌。毫不动摇坚持中国共产党的领导，确保人民政协事业发展的正确政治方向2019年9月20日，在中央政协工作会议暨庆祝中国人民政治协商会议成立70周年大会上，习近平总书记指出：“中国共产党的领导是包括各民主党派、各团体、各民族、各阶层、各界人士在内的全体中国人民的共同选择，是成立政协时的初心所在，是人民政协事业发展进步的根本保证。要把坚持党的领导贯穿到政协全部工作之中，切实落实党中央对人民政协工作的各项要求。”统一战线、武装斗争、党的建设，是中国共产党在革命斗争中形成的三大法宝。作为中国人民爱国统一战线的组织、中国共产党领导的多党合作和政治协商的重要机构，人民政协从诞生之日起，就毫不动摇坚持中国共产党的领导，确保人民政协事业发展的正确政治方向，与全国各族人民一道前进，与共和国一道成长，走过了光辉的历程。1949年9月21日至30日，中国人民政治协商会议第一届全体会议召开。会议代表全国各族人民意志，代行全国人民代表大会职权，通过了具有临时宪法性质的中国人民政治协商会议共同纲领和中国人民政治协商会议组织法、中华人民共和国中央人民政府组织法，作出关于国都、国旗、国歌、纪年的决议，选举产生政协全国委员会和中央人民政府委员会，标志着人民政协制度正式确立。75年来，人民政协自觉把中国共产党领导贯穿到协商议政、民主监督、凝聚共识、调查研究、团结联谊、自身建设等各项工作之中，把中国共产党主张转化为社会各界的广泛共识和自觉行动，确保中共中央决策部署在人民政协得到全面贯彻。进入新时代，以习近平同志为核心的中共中央全面加强对人民政协工作的领导，习近平总书记关于加强和改进人民政协工作的重要思想，为新时代人民政协担当新使命新任务指明了方向、提供了根本遵循。2018年3月4日，习近平总书记看望参加全国政协十三届一次会议的民盟、致公党、无党派人士、侨联界委员，并参加联组会时指出：“中国共产党领导的多党合作和政治协商制度作为我国一项基本政治制度，是中国共产党、中国人民和各民主党派、无党派人士的伟大政治创造，是从中国土壤中生长出来的新型政党制度。”“新型政党制度”的提出，彰显了坚定的制度自信。“人民政协要通过有效工作，努力成为坚持和加强党对各项工作领导的重要阵地、用党的创新理论团结教育引导各族各界代表人士的重要平台、在共同思想政治基础上化解矛盾和凝聚共识的重要渠道。”2019年9月20日，在喜迎新中国成立70周年、庆祝人民政协成立70周年之际，中共中央首次召开中央政协工作会议，习近平总书记发表重要讲话指出。2018年10月，中共中央办公厅印发了《关于加强新时代人民政协党的建设工作的若干意见》，指出中国共产党领导是中国特色社会主义最本质的特征，也是人民政协这一制度安排和政治组织最本质的特征。在政协各级组织和各项活动中，党是居于领导地位的，坚持中国共产党领导是人民政协必须恪守的根本政治原则。2019年10月7日，中共中央印发《关于新时代加强和改进人民政协工作的意见》，强调通过有效工作，使人民政协成为坚持和加强党对各项工作领导的重要阵地、用党的创新理论团结教育引导各族各界代表人士的重要平台、在共同思想政治基础上化解矛盾和凝聚共识的重要渠道。习近平总书记在中央政协工作会议上的重要讲话和党中央对新时代人民政协工作的一系列重要部署，为人民政协事业发展提供了根本政治保证，推动人民政协实践创新、理论创新、制度创新，为党和国家事业发展作出新贡献。发挥好人民政协专门协商机构作用，凝聚起实现中国梦的强大正能量在中国社会主义制度下，有事好商量，众人的事情由众人商量，找到全社会意愿和要求的最大公约数，是人民民主的真谛。协商民主是实现党的领导的重要方式，是我国社会主义民主政治的特有形式和独特优势。75年来，人民政协注重发挥专门协商机构作用，在协商中促进广泛团结、推进多党合作、实践人民民主，充分体现了我国社会主义民主有事多商量、遇事多商量、做事多商量的特点和优势。进入新时代，人民政协不断拓展协商内容、丰富协商形式，更加灵活、更为经常开展专题协商、对口协商、界别协商、提案办理协商，探索网络议政、远程协商、专家协商等新形式，提高协商实效。探索社会主义协商民主新形式、新方法、新路径，是人民政协工作的着力点。从2013年开始，全国政协专题议政性常委会会议由以往的一年1次增加到一年2次；2014年起，专题协商会从每年1次增加到2次；2013年10月起，双周协商座谈会“登场”。协商内容也由经济、社会拓展到政治、文化、生态等诸多领域。2024年9月2日，全国政协在京召开“推进高水平对外开放”专题协商会，近100位全国政协委员参加会议，全国政协外事委员会负责同志介绍专题调研情况，21位政协委员发言，围绕稳步扩大制度型开放、建设高标准市场体系、优化区域开放布局、推动产业链供应链国际合作、发展外贸新业态新模式、建设更高水平自贸试验区、推动共建“一带一路”高质量发展等协商建言。“要调动种粮农民和粮食主产区的积极性，强化政策支持，健全农业保险保障体系，构建现代化粮食产业体系，保障国家粮食安全。”2024年4月12日，十四届全国政协第十七次双周协商座谈会在京召开，围绕“调动种粮农民和粮食主产区的种粮积极性，夯实粮食安全的根基”协商议政，9位全国政协委员、2位种粮大户和1位专家代表先后发言，积极为保障国家粮食安全献计出力。大问题、小切口，形式活泼、气氛活跃。10年多来，双周协商座谈会每两周举办一次，每次邀请部分政协委员就关系国计民生的重大事项、人民群众关心的热点问题进行协商交流，如今已成为政协协商民主经常性平台和重要品牌。“双周协商座谈会围绕着增加协商密度、提高协商成效的新要求，在继承中创新，综合运用各种协商民主形式，以界别为基础、以专题为内容、以对口为纽带、以座谈为方式，创造了一种新型的协商民主形式。”中共中央党校原副校长李君如评价。2024年4月26日，全国政协召开远程协商会，围绕“深化人工智能多场景应用，提升现代产业高质量发展水平”协商议政。当天，10名全国政协委员和专家分别在全国政协机关和上海、浙江、安徽会场以及北京、广东等地，通过视频连线等方式发言，并同有关部委的负责同志就相关情况进行了协商交流。远程协商会是全国政协发挥互联网优势，有效提升协商议政质量，不断丰富协商形式的缩影。委员们表示，远程协商会更充分地发挥各种协商平台的独特优势和重要作用，参与面广、即时性强、互动性好，促进议题向更深更细发展，为人民政协工作注入了新的生机与活力。经过不断探索实践，全国政协逐步形成了以政协全体会议为龙头，以专题议政性常委会会议和专题协商会为重点，以双周协商座谈会、对口协商会、提案办理协商会等为常态的协商议政新格局。进入新时代，人民政协坚持把协商民主贯穿政治协商、民主监督、参政议政全过程，进一步搭建机制化、常态化的协商参与平台，通过加强制度化、规范化、程序化建设，逐步形成协商民主深入实践、各界活力充分释放、民意民智广泛汇聚、与党和政府畅通对接的生动局面，凝聚起实现中国梦的强大正能量。强化政协委员责任担当，积极建言献策，全面增强履职本领10月1日，是中华人民共和国的国庆日。而它的由来，与一份政协委员的提案有关。1949年10月9日，中国人民政治协商会议第一届全国委员会第一次会议正在紧张而热烈地进行。受因病未能参会的马叙伦委员委托，许广平委员提交了一份简短的提案：“请政府明定10月1日为中华人民共和国国庆日，以代替10月10日的旧国庆日。”最终，会议通过了《关于中华人民共和国国庆日的决议》。75年来，政协委员作为各党派团体和各族各界代表人士，代表各界群众参与国是、履行职责，通过提交提案、委员视察、专题调研、反映社情民意等多种方式积极建言献策，不断增强履职本领。2017年9月1日，第十二届全国人民代表大会常务委员会第二十九次会议通过了《中华人民共和国国歌法》，并于2017年10月1日起施行。而早在1995年，第八届全国政协委员蒋家祥等2人就在全国政协八届三次会议上提交了《关于尽快确定中华人民共和国国歌的法律地位并制定国歌法的提案》。2008年起，十一届、十二届全国政协委员于海连续10年提交提案，呼吁为中华人民共和国国歌立法。2017年3月，当即将离任全国政协委员的于海得知，自己关于国歌立法的提案得到批复，提上了议事日程时，不禁感慨万千。没有调查研究就没有发言权。人民政协将调查研究作为谋事之基、成事之道，组织委员认真调查研究、深入协商讨论，通过提高调研质量提升建言水平，为以中国式现代化全面推进中华民族伟大复兴贡献智慧和力量。2024年7月1日傍晚，黑龙江省大庆市人民政府2楼会议室里，全国政协调研组围绕“完善残疾人社会保障制度和关爱服务体系”开展专题调研，委员们利用晚间休息时间，与大庆市残联、专门协会、助残社会组织负责人和基层残疾人代表等恳谈交流，倾听他们的诉求与心声。重度残疾人托养照护需求迫切、特殊教育师资队伍相对薄弱、残疾人服务效能仍需进一步提高、残疾人精神文化生活日益多元……10余位来自基层一线的代表敞开心扉，道出了服务残疾人过程中的收获与困惑。委员们的及时回应，收获了在场的残疾人工作者以及残疾人代表的阵阵掌声。调研方式方法如何创新？十四届全国政协制定加强和改进调查研究工作的意见，加强对视察考察调研活动统筹安排，2024年共安排99项调研活动，实现对31个省（区、市）的全覆盖，避免了集中扎堆调研，减轻基层负担。建立由全国政协主席会议成员牵头、跨专门委员会专题研究制度。创新开展主席会议成员率队学习调研活动。有的调研组在调研过程中派出小分队，直接深入社区和乡村，到居民和农户家中听真话、察实情；有的调研组持续聚焦同一主题开展调研；有的调研组在调研期间关注群众急难愁盼问题，在当地群众心目中留下“人民政协为人民”“人民政协办实事”的良好形象；还有的调研注重形成合力，邀请有关部委、党派团体、智库机构等协同调研，邀请专家学者、当地委员随同调研，委托地方政协同步开展调研，广泛汇聚智慧和力量。2018年3月3日，全国政协会议首次在人民大会堂中央大厅北侧开设了“委员通道”，邀请来自多个界别的委员接受采访。6年来，每年都有数十名来自不同界别、不同领域的全国政协委员先后走上“委员通道”，与中外记者面对面，直接回答社会关心的话题。“委员通道”成为新时代全国政协全体会议的创新开放之举。2018年8月20日早晨，前来出席全国政协常委会会议的委员们发现，会场多处摆放着醒目的“二维码”，这是全国政协首次推出的委员移动履职平台。委员们表示，通过移动履职平台，实现了“思想永远在线、智慧时刻连线、联系永不断线”的随时协商。2024年6月26日，全国政协“委员科学讲堂”举办科普讲座，全国政协常委、中国科学院院士周忠和以“进化论与人类未来”为题作科普讲座，介绍进化生物学的发展与广泛影响，分享对进化论与人类演化的思考，并与现场观众互动交流。“委员科学讲堂”是全国政协扎实开展委员履职“服务为民”活动的一个缩影。2023年7月，政协第十四届全国委员会第六次主席会议决定在全国政协委员中开展委员履职“服务为民”活动。活动开展以来，全国政协委员关注民生福祉，回应群众期盼，充分发挥自身的优势和专长，为广大群众提供丰富多样的服务，让人民群众真切感受到政协履职的独特魅力与成效。教科卫体委员会赴新疆开展委员履职“服务为民”卫生“三下乡”活动，委员专家们下沉一线开展义诊会诊，造福广大基层群众；社会和法制委员会梳理出学者、在政法领域工作的委员情况和人数，作为法律宣传和免费法律咨询主体；外事委员会赴上海、广东等地开展服务外经贸企业高水平对外开放委员履职“服务为民”活动，帮助企业更好地“走出去”……在一场场活动中，委员们服务为民的方式越来越多元，服务为民的导向更加深入人心。新时代新气象，体现在人民政协每一次协商会议上、每一个履职活动中。新的思路理念，新的探索实践，新的精神风貌，新的担当作为，描绘着新时代人民政协守正创新的新画卷。75年和衷共济，留下辉煌足迹；75年奋进岁月，见证不变初心。团结奋进新征程，同心共筑中国梦，在以习近平同志为核心的中共中央坚强领导下，人民政协必将在新征程上谱写新的时代华章、创造新的历史伟业。`
var text2 = `9月20日，是北京环球度假区正式开园3周年。过去3年，“环球”已经迅速成长为北京乃至全国最火热的旅游目的地之一。游客从五湖四海而来，为了“环球”，又不止于“环球”。 在周边，民宿、足疗、餐馆……众多商家承接了“环球”带来的“泼天富贵”，将其外溢效应发挥到淋漓尽致。记者走访环球度假区周边区域，讲述3个普通经营者的“环球”3周年。 民宿 95%的客人是因为“环球” 距离环球度假区1.9公里，有一座通州区典型的平原村——唐大庄村。这里原本以观赏鱼业闻名，如今，摇身一变，成为“环球最近民宿村”。 “当时在周边考察，穿过一个桥洞，然后豁然开朗，三片林地包围着一个村庄，有种世外桃源的感觉。”梁筱娟，曾是少见的计算机编程专业“理工女”，后来做工程，现在是暖唐民宿主理人。“刚开始只是想当个副业来做，没想到规模越做越大。” 家在通州区的梁筱娟，早就有做民宿的想法。2021年，在环球度假区正式开业前夕，通州区组织多部门联合召开规范短租住房经营管理工作部署会。 “我觉得时机成熟了。”在充分考察后，梁筱娟与正在积极寻求转型的唐大庄村不谋而合，“在村委会的民宿项目接待室，对着沙盘，听村领导班子一介绍，我心想‘行，就是这儿了’。” 曾经的唐大庄村，聚集着低廉的瓦片经济，一度被定义为集体经济薄弱村。2022年1月，梁筱娟打造的暖唐民宿开张。如今，全村145个院落中，有约100个出租做民宿，约80个正式签约。其中，约30个院落属于暖唐民宿，共130间客房。暖唐也成为北京民宿客房单体最大的本土企业。“村民们从低廉的瓦片经济中走出来，现在的房屋租金是原来的10倍。” 走在现在的唐大庄村，魔法主题餐厅，狮院主题民宿……各式各样的“环球”元素点缀在乡间公路两侧。 在狮院主题民宿，室内装饰以红色为主，随处可见以雄狮为核心的徽章。梁筱娟说：“很多游客是资深粉丝，很多人就是奔着这个来的。” 在魔法主题餐厅，魔法酸汤鱼、魔法辣子鸡，是菜单上的热门选择。“我们有往返大巴，游客下大巴可以直接进餐厅。用完餐，电瓶车摆渡回民宿。如果太累了，游客可以在民宿里点餐，我们送餐。”  在刚刚过去的暑假，暖唐民宿基本满房，其中有95%是“环球”游客。“全年大概能接待65000人，入住率在70%左右。这在民宿里，已经是相当高了。”  目前，唐大庄村，正在进行整体环境提升。梁筱娟在积极寻求配套商业合作，希望未来，游客们除了住进来，还能留几天，有得玩、有得逛。 足疗 游客穿着魔法袍来放松 “暴晒、暴走，一整天，脚都要废了，找个地方做足疗，疲惫逐渐被愉悦替代。”“环球走了一天，酒店旁边就有足疗店，太舒服了。”“日行两万步，逛完环球，打车起步价，足疗太值了。” 在社交媒体上，很多游客的笔记里，将“环球”与足疗联系到一起。在环球度假区周边5公里范围内，如今云集着约150家足疗按摩店。 “芭娜娜谢谢你足道”地处九棵树，与环球度假区直线距离4公里。创始人于桂梅，1997年来到北京，2005年开始创业做足疗。自“环球”开业后，她才见到穿着魔法袍来放松的客人。 “我们这家店是去年5月28日开张的，开业没多久，就遇上了暑期旅游高峰。”于桂梅说，以往，通州区足疗店的主要客群是周边居民和商旅客人，“中年人比较多，游客也有，年轻的游客不多。”  “环球”带来了巨大改变，穿着魔法袍来的，是热衷“特种兵”式旅游的年轻游客。一大早进园，狂走几万步，等出园那一刻，已经精疲力竭。打开手机，查看攻略，是搜索附近足疗店的时候了。于桂梅的足疗店，有柔和的灯光、舒适的躺椅、美味的小吃。游客们可以按摩放松，品茶，看电影，尝小吃。70分钟的足疗，或者90分钟的全身按摩，帮游客卸下一身的疲惫。 “有些游客是第一次来北京，第一站就是‘环球’，逛完‘环球’，第一次消费，是我们足疗店。”于桂梅说，每次有游客反馈“舒服、温暖”，都是自己最满足的时刻，“我能理解游客初次到北京的心情，就跟当年我来‘北漂’一样。在一个陌生的大城市，最难得就是找到像家一样的温暖。” 足疗生意火爆，竞争对手变得更多，于桂梅倒是很乐意接受竞争。“我自己在这个行业20年，知道服务是最核心的竞争力。从客人进店开始，接待、进屋、上茶水水果、技师按摩……每个环节，都尽力做到专业贴心。”当游客们说出：“下次来‘环球’，还来你家店。”于桂梅知道，努力没白费。 如今，儿子已经开始帮她打理店面。在通州区永顺镇的新店面，也正在装修当中。“那边民宿和酒店非常多，新的足疗店，可以更好地为外地游客服务。” 餐馆 游客进门就问“有烤鸭吗” 距离环球度假区外墙仅500多米，有一片围绕居民区的商业配套。紫光园·烤鸭（加州小镇店）店长杨光，会在休息的时候，站在路口，远眺“环球”：“能看见过山车。” 餐饮做了快20年，前15年是掌勺做厨师，来北京4年，开始当店长成为管理者。2021年，杨光来到这家紫光园，管理刚刚开业的新店。“我们店，包括周围其他一些商业，酒店、足疗什么的，差不多都是同一时间开业，就在‘环球’开园后不久。” 周围全是居民区，饭店的主要客群也是周边居民。但每年一到暑期和节假日，客群就开始变化，穿着魔法袍的游客变多了。“尤其是去年和今年两个暑假，非常明显。”游客到店多，来自周边酒店的外卖订单也多。 有游客进门就问：“有烤鸭吗？”这家店面积不大，但特色鲜明，店名主打烤鸭，吸引想尝鲜北京烤鸭的外地游客。“现在游客都是打开手机App，查看点评。咱们这边距离有优势，有北京特色，而且主打便民亲民惠民，价格实惠。” 店隔壁就是一家酒店，游客从住处下楼，走不了几步，就能进店品尝烤鸭。为了满足游客们对烤鸭和北京菜的好奇心，店里还推出了套餐。烤鸭、椒盐鸭架、醋熘木须、酸辣汤，两到三人的量，现在只要149.9元。 作为店长，杨光每天都早早到店。周五、六、日，甚至跟着做早点的同事一起，在凌晨2点到店。“咱们店都是现做的，早餐价格也不贵，5元吃饱，8元吃好。”忙完早餐，收拾店面，差不多快10点了。 接下来就要准备烤鸭。为了11点就能让午餐客人吃到烤鸭，鸭胚要提前大约四五十分钟就进炉烤制。走进烤鸭厨房，炉火正旺。“今天准备了多少只？”“（午餐）15只。”杨光对烤鸭师傅点点头。  这是平日里正常的预备量，晚餐比午餐多，一天下来差不多能卖约40只烤鸭。在暑期，一天能卖约60只烤鸭。 环球度假区开园3年，这家紫光园也开业3年，“环球”的外溢效应，伴随着这家小店的整个成长历程。“一晃就3年了，时间还真挺快。服务好周边居民，给游客留下好印象，这3年，我觉得我们做到了。”国庆假期将至，又一个旅游高峰，杨光说，店里已经做好充足准备，热腾腾的烤鸭，一出炉就上桌。 数说外溢 2023年，“环球”商圈客流量约1600万人次，通州区规模以上文化、体育和娱乐业企业实现收入96.8亿元，2013年以来年均增长60.7%。 2021年，在“环球”有力带动下，通州区规模以上文化、体育和娱乐业收入同比增长367.4%，住宿业收入增长122.6%。 2021年，“环球”开园首月，周边5公里内的按摩足疗订单量上涨74.6%。 2021年10月到2022年6月，在“环球”带动下，通州区接待游客同比增长46.7%，增速全市排名第一。“环球”周边、梨园、北苑区域游客游览热度较高，承接游客占比90.9%。 2021年10月到2022年6月，在“环球”带动下，通州区实现旅游收入59.9亿元，排名全市第七，同比上升4位；增速44.7%，排名全市第一。 2022年暑期、“十一”相关监测数据显示，“环球”共接待游客104.1万人次，其中通州区承接了40.3万人次的环球影城外溢游客，占比达到38.7%。 “环球”开业前，周边有餐饮、购物、酒店（含民宿）、休闲娱乐、运动健身等门店1800余家，“环球”开业后，增加至2100余家。分类型看，新增门店主要集中在餐饮、休闲娱乐和酒店（含民宿）领域，分别新增102家、76家和67家。 （数据来源：北京市统计局、北京市政府新闻办、北京市文旅局、美团、大众点评、去哪儿网等）`
var text3 = `时隔一天，黎巴嫩多地再次发生通信设备爆炸事件。两天之内，类似事件已造成逾30人死亡，4000多人受伤，遇难者中包括多名儿童和医护人员。 黎巴嫩真主党指责以色列对事件负有“全部责任”，并誓言报复；国际社会纷纷谴责爆炸事件，普遍认为是以色列策划了一系列袭击。以色列则照例采取“不承认不否认”的态度，至今没有回应。 一系列袭击看似针对黎巴嫩真主党成员，但发生在闹市的爆炸导致不少无辜平民被殃及，伤亡惨重。全球多个人权组织和活动人士直言，以色列此举堪称国家恐怖主义。 另一方面，此次爆炸袭击所使用的手段和规模可谓前所未有，也敲响了供应链安全警钟。有美媒直言，从寻呼机到对讲机的爆炸，都预示着未来战争的战线可能无限延伸，日用品也不可信，而供应链战争将会持续下去。 “从未在一天之内摘除那么多眼球” 以色列和黎巴嫩真主党的矛盾由来已久。去年10月新一轮巴以冲突爆发后，真主党频繁袭击以色列北部军事目标，以策应巴勒斯坦伊斯兰抵抗运动（哈马斯），以军则回以空袭和炮击。今年7月，真主党高级指挥官福阿德·舒库尔（Fouad Chokr）遭以色列空袭身亡，进一步点燃双方战火。 然而，此次黎巴嫩通讯设备的大规模爆炸还是震惊了国际社会。 为躲避以色列定位，真主党武装人员一直在使用寻呼机作为一种低技术通信手段。消息人士表示，以色列情报机构摩萨德数月前就在真主党订购的5000台寻呼机内安放了爆炸物，而18日爆炸的那批对讲机与购买寻呼机的时间大致相同。此外，一些报道提及，黎巴嫩部分太阳能设备和汽车中的电池当天也被引爆。 17日的报道和视频显示，爆炸发生在超市、水果摊、街道等公共场所，至少导致两名儿童和四名医护人员死亡。数千名伤者中，近三分之二的人需要对面部、眼睛或手部进行手术，许多人不得不截肢。  当地时间2024年9月18日，黎巴嫩提尔古城，受伤人员正在接受治疗。 视觉中国 爆炸发生时，黎巴嫩作家、前驻约旦大使特雷西·夏蒙（Tracy Chamoun）正驾车行驶在黎巴嫩首都贝鲁特南郊的一座立交桥上。突然，她看到有人躺在地上。 “当他们开始争先恐后地把这些人送往医院时，一切都变得一团糟。汽车被推到一边，摩托车和汽车载着满身是血的人开了过来。”夏蒙对英国广播公司（BCC）描述道，“我看到其中一名伤员眼睛被打瞎，另一个人的半边脸被撕掉了。” 夏蒙表示，寻呼机在被引爆前会发出声响，吸引人们把它拿出来查看。这种说法和此前伊朗消息人士对《纽约时报》所描述的类似，即一些受害者查看寻呼机时，眼睛和眼部被炸伤。伊朗驻黎巴嫩大使阿马尼也在此次爆炸中被炸伤眼睛。 贝鲁特黎巴嫩山医院的一名眼科医生表示，过去的24小时就如同一场“噩梦”，超过60%至70%的伤者被摘除了一只眼睛。“有些病人不得不摘除双眼，这简直是杀了我，在我过去25年的执业生涯中，从来没有像昨天那样摘除那么多只眼睛。” 不少评论指出，此次袭击看似针对黎巴嫩真主党成员，实则是无差别攻击。 据半岛电视台报道，总部位于华盛顿的人权组织“现在为阿拉伯世界民主”主任惠特森（Sarah Leah Whitson）表示，“你不应该在平民可能捡起和使用的物品上设置陷阱”，“这正是我们在黎巴嫩看到如此惨烈灾难的原因。”惠特森指出，大量人员的伤亡表明，这些袭击“本质上是无差别的”。 美国人权律师胡瓦伊达·阿拉夫（Huwaida Arraf）认为，以色列此举违反了国际人道法和联合国《禁止或限制使用地雷、诱杀装置和其他装置的议定书》 在阿拉夫看来，只有采取措施保护平民，并确保爆炸只击中合法的军事目标，这些袭击才能被认为是合法的。然而，黎巴嫩的这些爆炸在发生之前没有任何预警，且发生在公共场合，这“事实上符合对国家恐怖主义”的定义。 英国工党前领袖科尔宾（Jeremy Corbyn）创立的“和平与正义计划”18日发声明，直言以色列17日的袭击堪称“国家恐怖主义”，必须受到国际社会的谴责，该组织还呼吁对以色列实施“重大制裁”。  “和平与正义计划”X账号截图 上海外国语大学中东研究所教授、中国中东学会副会长刘中民19日对观察者网表示，黎巴嫩真主党成员因通讯设备爆炸遭到打击并造成人员伤亡，体现了新一轮巴以冲突外溢的新形式、新特点、新挑战，导致战争与民事生活之间的界限更加模糊，尤其是其暴力威慑效应更加直接。 “这在某种程度上确实带有‘国家恐怖主义’的色彩，也凸显了以色列军事、情报和安全机构在打击异己方面无所不用其极的极端性。”刘中民表示。 据报道，联合国安理会将在当地时间9月20日下午就近期黎巴嫩多地发生的通信设备爆炸事件举行紧急会议。此前联合国秘书长古特雷斯在一份声明中对黎巴嫩所发生的事件表示震惊，呼吁各方保持最大克制，以避免事态进一步升级。 联合国人权事务高级专员蒂尔克谴责道，这次袭击事件对平民造成的影响令人无法接受，所引发的恐惧和恐怖是深远的。 “同时将数千人作为攻击目标，无论是平民还是武装团体成员，且在攻击时不知道谁拥有目标装置、其位置和周围环境，这违反了国际人权法，并在适用范围内违反了国际人道主义法。”蒂尔克表示呼吁彻底调查，追究下令并实施袭击的人员的责任。 敲响国际供应链警钟 西北大学中东研究所副教授王晋对观察者网表示，过去以色列会在某个特定人物的手机中植入炸弹等装置，实施其所谓的“定点清除”，但如此大规模的通信设备爆炸，此前还没有先例。 在王晋看来，此举一是为了打乱真主党的组织架构，“使用BP机和对讲机进行通讯的多为真主党中下层的骨干人员，因此这次行动在较大程度上损坏了真主党的组织动员能力，短期内给该组织造成了较大的负面影响。” 第二则是给真主党武装人员制造巨大的心理恐慌，“他们不知道下一次爆炸何时会发生，身边其他物件是不是会发生爆炸。” 王晋表示，真主党认定是以色列所为，势必会对其发动大规模攻击。由于真主党的架构在此次事件中受到较大影响，报复措施可能需要等一段时间，“但无论如何，真主党的报复是迟早会来的，只是时间问题。而以色列也会继续向黎境内真主党目标发动袭击。” 王晋提到，在他国发动这样的攻击，除了进一步破坏地区与国际局势的和平稳定，对于国际贸易的正常秩序也是巨大冲击。 有关爆炸物如何被置入通讯设备的细节，目前还众说纷纭。但在消息人士和分析人士看来，以色列显然已经渗透了供应链。有黎巴嫩高级消息人士表示，以色列情报机构已经“在生产阶段”对这些设备进行了修改。  发生爆炸的寻呼机 美联社援引社交媒体图片 “这看起来可能是史上最大规模的实体供应链攻击。”美国智库“西尔维拉多政策加速器”（Silverado Policy Accelerator）主席阿尔佩洛维奇（Dmitri Alperovitch）对《华盛顿邮报》表示。 据报道，17日发生爆炸的寻呼机为台湾地区品牌“金阿波罗”。该公司18日紧急切割，表示涉事产品由一家名为BAC的欧洲公司生产和销售，“金阿波罗”只是提供品牌商标授权，BAC公司地址位于匈牙利首都布达佩斯。 然而，匈牙利政府18日也发声明澄清，表示黎巴嫩大规模爆炸中使用的寻呼机设备从未出现在匈牙利，BAC公司只是一家贸易中介公司，在匈牙利无制造厂。 《纽约时报》18日披露的最新细节显示，以情报部门利用了真主党转向使用寻呼机的倡议，在很早之前就成立了至少3家空壳公司，伪装成国际寻呼机生产商进行生产销售，将大量掺杂了炸药的寻呼机送进黎巴嫩。这些寻呼机堪称“以色列制造的现代特洛伊木马”。 至于18日发生爆炸的对讲机，其制造商被指是一家总部位于日本的无线电通信设备公司。而后者回应称，发生爆炸的对讲机疑似为山寨仿造品，企业正在调查。 无论是哪个环节出的问题，本周黎巴嫩所发生的一系列事件都为全球供应链安全敲响了警钟。 “从寻呼机、对讲机到太阳能系统爆炸的报道，都预示着未来战争的战线可能无限延伸，日用品也不可信。”美国Axios新闻网评论道。 报道认为，未来供应链战争将会持续下去，也势必引发对供应链保障的新要求，因为对外国零部件的依赖会削弱该国的军事实力。 澳大利亚广播公司（ABC）称，抛开地缘政治问题和对中东未来的影响不谈，此事表明国家安全是如何通过看似平凡的技术而受到损害的，它突出了供应链脆弱带来的巨大风险。  当地时间9月18日，黎巴嫩多地再次发生通信设备爆炸事件。 视觉中国 刘中民也指出，以色列对供应链的渗透是很令人担心的事情，这导致军事与经济民事、冲突与和平之间的界限变得模糊，也成为在技术变革时代军事和安全领域的新灰色地带。 刘中民看来，黎巴嫩所发生的爆炸事件，应该不会对中东的供应链构成系统性威胁，但无疑使供应链安全变得更加脆弱，“如被恐怖组织和极端组织使用，则危害更大。” 此外，刘中民认为，未来以色列对伊朗、哈马斯、黎巴嫩真主党、也门胡塞武装等反以力量有进一步使用此类袭击手法的可能，但在范围和规模上会高度谨慎。 “以色列知道，他们可以无视美国” 从加沙地带、约旦河西岸到黎巴嫩，从“定点清除”到无差别攻击，以色列似乎有着“打不完”的仗，这也给相关地区民众带来了巨大灾难。 本轮巴以冲突爆发即将一年，据巴勒斯坦有关部门报道，以色列所展开的一系列军事行动已导致加沙地带逾4.1万名巴勒斯坦人死亡，近10万人受伤；约旦河西岸有超过650名巴勒斯坦人死于以军或犹太定居者之手。另据半岛新闻网报道，黎巴嫩有近600人死于以色列的战火。 “中东正在发生危险的战争升级，而这场战争的起因是以色列。内塔尼亚胡不满足于对加沙地带的种族灭绝性袭击，他决心在整个地区发动战争。”英国政治活动家、停止战争联盟林赛·赫尔曼（Lindsey German）说，她同样将黎巴嫩近期的爆炸事件称为“无差别的国家恐怖主义行为”。 半岛电视台专栏作家贝伦·费尔南德斯（Belen Fernandez）则在其评论文章中直指以色列是“一个专门以打击恐怖主义为名、恐吓特定阿拉伯平民的国家”。 “尽管周二袭击的表面目标是手持寻呼机的真主党成员，但袭击者在实施袭击时完全知道后果将是无差别的，会造成大量平民伤亡。但这正是恐怖主义的核心，难道不是吗？”  当地时间2024年9月18日，黎巴嫩贝鲁特南部地区举行的葬礼上，妇女们举着死者的照片。 视觉中国 对于以色列再次激化地区紧张局势的行为，其盟友美国也再次表示对此事不知情。Axios新闻网18日援引三名美国官员的话说，黎巴嫩寻呼机爆炸事件发生前几分钟，以色列国防部长加兰特曾致电美国防长奥斯汀，称以色列将在黎展开行动，但未透露具体细节。 美国国务卿布林肯回应称，美国对此事并不知情，也未参与其中。 美国阿拉伯反歧视委员会（ADC）9月17日发声明，强烈谴责以色列实施“无差别恐怖主义行为”的同时也强调，以色列不顾一切地寻求战争，正是得益于拜登政府的纵容和无条件支持。 “以色列领导人知道，他们可以无视美国的建议而不必担心后果，他们知道，美国纳税人将继续提供武器和财政支持。”声明称。 王晋认为，美国或许对以色列此举并不完全知情，“或者说美国没有办法完全影响和控制以色列的具体决策。” 当地时间9月18日，联合国大会在讨论巴以问题的第十次紧急特别会议上，投票通过关于以色列在巴勒斯坦被占领土问题的咨询意见的决议草案，要求以色列在决议通过后12个月内结束其在巴勒斯坦被占领土的非法存在。 表决中，124票赞成，美国、以色列等14票反对，43票弃权。中国投下赞成票。 中国常驻联合国代表傅聪在9月17日的发言中表示，中方敦促以色列倾听国际社会的强烈呼声，立即结束对巴勒斯坦领土的非法占领。“结束占领不是选择，而是以色列的法律义务。”傅聪说。 来源|观察者网`
var text4 = `急诊工作中经常会遇见吐血的病人，医生在做检查前都会询问一句： 您感觉血是恶心呕吐出来的，还是咳嗽出来的？ 部分患者听到这个问题会很疑惑，血就是从嘴里吐出来的呀？还有区分？ 其实人体与外界相通的腔道分为消化道和呼吸道，二者有一个共同的部位 咽，因此消化道与呼吸道中任何一个部位的出血都有可能从口腔流出。 临床上消化道出血是指食道以下的消化系统出血，可有呕血；呼吸道出血是指喉及喉部以下呼吸器官出血，又称为咯血。无论呕血还是咯血，疾病风险程度都比较高，二者的治疗方案有本质上区别，因此临床医生需要根据第一个问题的答案来初步划分是消化道出血还是呼吸道出血。 急性上消化道出血的处理 上消化道出血是指十二指肠悬韧带以上的食管、胃、十二指肠、上段空肠以及胰管和胆管的出血。 '消化道简易示意图' 急性上消化道出血主要表现为呕血和（或）黑便。当出血量在50ml，可出现黑便；胃内积血量在250ml以上，可引起呕血，由于胃酸的作用，呕吐物颜色多为咖啡色；一次出血量 400ml，可出现头晕、乏力、心悸等症状，此时因为出血量多、在胃内停留时间短，呕吐物颜色呈现鲜红色或暗红色，可伴有血块。 第一步处理：禁食禁水，卧位休息。呕吐频繁时头偏向一侧，避免呕血时血液吸入引起窒息，保持呼吸道通畅，必要时吸氧，严密监测患者血压心率变化。尽早补液治疗，可以行急诊胃镜检查，明确出血部位和性质，以及镜下止血。若患者出血量较大或无法配合胃镜检查，可以进行介入下血管栓塞。 下呼吸道出血的处理 鼻腔、咽、喉、气管、主支气管和肺组成了呼吸道。喉及喉部以下呼吸器官称为下呼吸道，此处出血又叫咯血。每日咯血量在100ml以内为小量咯血，500ml以上或一次咯血100~500ml为大咯血。 由于引起咯血的疾病不同，出血颜色也有区别，如铁锈色血痰见于肺炎链球菌性肺炎，鲜红色痰见于支气管扩张，浆液性粉红色泡沫痰则多发于急性左心衰竭。 第一步处理：保持镇静，及时就医。吸氧、监护、止血、输液等对症治疗的同时，完善检查化验寻找病因。需要重点说明的一点，如果咳嗽不剧烈，不需要使用镇静止咳药物，以免抑制咳嗽中枢，使得血液淤积气道，引起窒息。 口腔出血的处理 血液经口腔流出的原因除了前面讲的呕血和咯血，还可能是口腔出血。 口腔出血常由于创伤、牙龈炎、牙周炎等所致，也可能是由于血液疾病等引起口腔出血。曾经有一名患者，满口鲜血地来就诊，通过询问病史、查体、检验等一系列检查，排出了消化道出血、呼吸道出血，也排除了血液疾病继发的出血表现，而且各种经验性治疗效果不理想。正当医生百思不得其解的时候，患者说了一句话，让她茅塞顿开。患者说，感觉每次出血都是从嘴里涌出来的。赶紧请口腔科会诊，发现患者有严重的牙周疾病，出血的原因正是牙周炎引起的。 第一步处理：压迫止血。如果患者在家中可判断是口腔出血，那么就紧紧咬住纱布或者毛巾，进行压迫止血。如果出血止不住，及时到医院处理。 鼻出血的处理 鼻出血也可以从口腔中流出。鼻出血常见的出血部位在鼻中隔前下方和后鼻孔外侧壁。婴幼儿和青少年由于挖鼻、鼻腔异物、外伤、鼻炎等原因发生鼻出血，最常见的部位是鼻中隔前下方；中老年患者的高血压、动脉硬化是鼻出血的主要原因，多见于后鼻孔外侧壁，出血一般较为凶猛，不易止血，出血常迅速流入咽部，从口中吐出，因而危险性更大。前者可在家中处理，后者赶紧送医救治。 第一步处理：指压法止血。如果出血少量且出血在鼻腔前部的患者，可在家中采用指压法止血。用手指捏紧双侧鼻翼或将出血侧鼻翼压向鼻中隔约10～15分钟，注意头部呈轻微前倾的位置，切勿抬头，因为抬头容易使血液经咽部流入气道，有引发窒息的危险。若这样处理后口中仍吐出鲜血，也可用手指横行按压上唇部位，同时冷敷前额和后颈部。 如果出血量没有减少，需立即就诊，医生会采用喷药、填塞、栓塞等方法帮助止血，同时控制血压，治疗原发疾病。 经口出血是临床工作中常见的症状，需要医生胆大、心细，也需要患者积极配合，共同查找出血真相，遏制疾病发展，尽早解除患者痛苦。 作者：十院科普官、上海市第十人民医院急诊医学科 罗璧君 主治医师`
var text = text1 + text2 + text3 + text4

var exp = "/"
var find = map[string]int{}
var dict = []string{"中国梦", "人民政协", "紫光园", "环球度假区", "黎巴嫩真主党", "梁筱娟"}
var dictStr = exp + strings.Join(dict, exp) + exp
var sg Segment

func TestCut(t *testing.T) {
	sg = &Jieba{}
	Cut(t)

	sg = &Gse{}
	Cut(t)
}

func Cut(t *testing.T) {
	sg.Init()

	startTime := gtime.Now().TimestampMilli()
	err := sg.LoadDictWords(dict)
	if err != nil {
		t.Error(err)
		return
	}

	healthDict := dicts.LoadHealthDict()
	err = sg.LoadDictWords(healthDict)
	if err != nil {
		t.Error(err)
		return
	}

	token := sg.Cut(text)
	for _, token := range token.Words {
		if strings.Index(dictStr, exp+token+exp) >= 0 {
			find[token] += 1
		}
	}
	t.Log("load dict use time (ms)", gtime.Now().TimestampMilli()-startTime)

	if len(find) <= 0 {
		t.Error("No dict words founds")
	}

	t.Log(find)
	t.Log(strings.Join(token.Words, exp))

	startTime = gtime.Now().TimestampMilli()
	t.Log(sg.Pos(text))
	t.Log(sg.Ner(text))
	t.Log("cut use time (ms)", gtime.Now().TimestampMilli()-startTime)

	var ners = gset.NewStrSet()
	for _, ner := range sg.Ner(text) {
		if gstr.LenRune(ner.Word) > 3 {
			ners.Add(ner.Word)
		}
	}
	t.Log(ners)
}

func BenchmarkGse(b *testing.B) {
	sg = &Gse{}
	sg.Init()

	err := sg.LoadDictWords(dict)
	if err != nil {
		b.Error(err)
		return
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sg.Cut(text)
		}
	})
}

func BenchmarkJieba(b *testing.B) {
	sg = &Jieba{}
	sg.Init()

	err := sg.LoadDictWords(dict)
	if err != nil {
		b.Error(err)
		return
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sg.Cut(text)
			sg.Pos(text)
		}
	})
}
