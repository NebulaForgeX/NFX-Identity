import { memo } from "react";
import { Button, Flex } from "@radix-ui/themes";
import { FileQuestion } from "lucide-react";

import { EmptyState } from "@/components";
import { routerEventEmitter } from "@/events/router";
import { PageFrame } from "@/layouts";
import { ROUTES } from "@/navigations";

const NotFoundPage = memo(() => {
  return (
    <PageFrame>
      <EmptyState
        icon={FileQuestion}
        title="Page Not Found"
        description="The page might have been moved or deleted. Please check the URL or go back to the homepage."
        action={
          <Flex gap="3">
            <Button variant="soft" onClick={() => routerEventEmitter.navigateBack()}>
              Go Back
            </Button>
            <Button onClick={() => routerEventEmitter.navigate({ to: ROUTES.USER_OVERVIEW })}>Go to Homepage</Button>
          </Flex>
        }
      />
    </PageFrame>
  );
});

NotFoundPage.displayName = "NotFoundPage";
export default NotFoundPage;
