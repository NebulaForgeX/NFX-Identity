import { memo } from "react";
import { Card, Flex, Heading, Text } from "@radix-ui/themes";
import { ExternalLink, FileText, History, Home, Info, Shield } from "lucide-react";

const ResourceLinks = memo(() => {
  const resourceLinks = [
    { title: "项目主页", description: "访问项目主页", icon: Home, url: "#" },
    { title: "更新日志", description: "最新更新和功能", icon: History, url: "#" },
    { title: "关于我们", description: "关于项目的信息", icon: Info, url: "#" },
    { title: "服务条款", description: "用户条款和条件", icon: FileText, url: "#" },
    { title: "隐私政策", description: "我们如何保护您的数据", icon: Shield, url: "#" },
  ];

  return (
    <Flex direction="column" gap="3">
      <Heading size="4">资源与政策</Heading>
      <Text size="2" color="gray">
        了解更多关于平台和政策
      </Text>
      <Flex gap="3" wrap="wrap">
        {resourceLinks.map((link) => (
          <Card
            key={link.title}
            size="2"
            asChild
            style={{ flex: "1 1 200px", cursor: "pointer" }}
          >
            <button type="button" onClick={() => window.open(link.url, "_blank", "noopener,noreferrer")}>
              <Flex direction="column" gap="2" align="start">
                <link.icon size={18} />
                <Heading size="3">{link.title}</Heading>
                <Text size="1" color="gray">
                  {link.description}
                </Text>
                <ExternalLink size={14} />
              </Flex>
            </button>
          </Card>
        ))}
      </Flex>
    </Flex>
  );
});

ResourceLinks.displayName = "ResourceLinks";
export default ResourceLinks;
