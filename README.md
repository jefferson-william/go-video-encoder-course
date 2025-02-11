# go-video-encoder-course

Converter vídeos para MP4 de acordo com melhor formato de vídeo para playlist.

## TODO

- [] Receber uma mensagem via RabbitMQ informando qual o vídeo deve ser convertido
- [] Fazer download do vídeo
- [] Fragmentar o video
- [] Converter o vídeo para mpeg-dash
- [] Fazer o upload do vídeo
- [] Enviar uma notificação via fila com as informações do vídeo convertido ou informando erro na conversão
- [] Em caso de erro, enviadar a mensagem original a uma Dead Letter Exchange

## Input

```json
{
  "resource_id": "my-resource-id-can-be-a-uuid-type",
  "file_path": "convite.mp4"
}
```

## Output

```json
{
  "id": "<uuid>",
  "output_bucket_path": "codeeducationtest",
  "status": "COMPLETED",
  "video": {
    "encoded_video_folder": "<uuid>",
    "resource_id": "<uuid>",
    "file_path": "convite.mp4"
  },
  "Error": "",
  "created_at": "2020-05-27T19:43:34.850479-04:00",
  "updated_at": "2020-05-27T19:43:38.081754-04:00"
}
```

## Error

```json
{
  "message": {
    "resource_id": "<uuid>",
    "file_path": "convite.mp4"
  },
  "error": "Motivo do erro"
}
```