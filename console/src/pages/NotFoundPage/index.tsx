import { memo } from "react";
import { Button, Flex } from "@radix-ui/themes";
import { FileQuestion } from "lucide-react";
import { PageFrame } from "nfx-ui/layouts";
import { EmptyState } from "nfx-ui/components";

import { TruckLoading } from "@/animations";
import { routerEventEmitter, routerEvents } from "@/events/router";

const NotFoundPage = memo(() => {
  return (
    <PageFrame>
      <EmptyState
        icon={FileQuestion}
        title="Page Not Found"
        description="The page might have been moved or deleted. Please check the URL or go back to the homepage."
        action={
          <Flex direction="column" align="center" gap="4">
            <TruckLoading size="medium" />
            <Flex gap="3">
              <Button variant="soft" onClick={() => routerEventEmitter.emit(routerEvents.NAVIGATE_BACK)}>
                Go Back
              </Button>
              <Button onClick={() => routerEventEmitter.emit(routerEvents.NAVIGATE_TO_HOME)}>Go to Homepage</Button>
            </Flex>
          </Flex>
        }
      />
    </PageFrame>
  );
});

NotFoundPage.displayName = "NotFoundPage";
export default NotFoundPage;
