/** Browser PUT to a MinIO/S3 presigned URL. Does not go through axios. */

export async function putToPresignedUrl(uploadUrl: string, body: Blob, contentType: string): Promise<void> {
  let res: Response;
  try {
    res = await fetch(uploadUrl, {
      method: "PUT",
      body,
      headers: { "Content-Type": contentType },
    });
  } catch (err) {
    if (err instanceof TypeError) {
      throw new Error("MINIO_UNREACHABLE");
    }
    throw err;
  }
  if (!res.ok) {
    throw new Error(`MINIO_HTTP_${res.status}`);
  }
}

export function minioUploadMessage(error: unknown, network: string, fallback: string): string {
  if (error instanceof Error && error.message === "MINIO_UNREACHABLE") return network;
  if (error instanceof Error && error.message.startsWith("MINIO_HTTP_")) {
    return `${fallback} (${error.message.slice("MINIO_HTTP_".length)})`;
  }
  return fallback;
}
