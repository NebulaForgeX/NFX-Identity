import { ExternalLinkIcon, FileDescriptionIcon, HistoryCircleIcon, HomeIcon, InfoCircleIcon, ShieldCheck } from "nfx-ui/icons";
import { memo } from "react";
import { Card, Flex, Heading, Text } from "@radix-ui/themes";
const ResourceLinks = memo(() => {
  const resourceLinks = [
    { title: "项目主页", description: "访问项目主页", icon: HomeIcon, url: "#" },
    { title: "更新日志", description: "最新更新和功能", icon: HistoryCircleIcon, url: "#" },
    { title: "关于我们", description: "关于项目的信息", icon: InfoCircleIcon, url: "#" },
    { title: "服务条款", description: "用户条款和条件", icon: FileDescriptionIcon, url: "#" },
    { title: "隐私政策", description: "我们如何保护您的数据", icon: ShieldCheck, url: "#" },
  ];

  return (
    <Flex direction="column" gap="3">
      <Heading size="4">资源与政策</Heading>
      <Text size="2" color="gray">
        了解更多关于平台和政策
      </Text>
      <Flex gap="3" wrap="wrap">
        {resourceLinks.map((link) => {
          const Icon = link.icon;
          return (
          <Card
            key={link.title}
            size="2"
            asChild
            style={{ flex: "1 1 200px", cursor: "pointer" }}
          >
            <button type="button" onClick={() => window.open(link.url, "_blank", "noopener,noreferrer")}>
              <Flex direction="column" gap="2" align="start">
                <Icon size={18} />
                <Heading size="3">{link.title}</Heading>
                <Text size="1" color="gray">
                  {link.description}
                </Text>
                <ExternalLinkIcon size={14} />
              </Flex>
            </button>
          </Card>
          );
        })}
      </Flex>
    </Flex>
  );
});

ResourceLinks.displayName = "ResourceLinks";
export default ResourceLinks;
