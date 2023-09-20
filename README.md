# sdmodels-json

This project uses the Go programming language for categorizing different models like `Lora`, `Checkpoint`, `Vae`, and `Embedding`. It reads the data from the file, categorizes them accordingly and prints them.

## Usage

To run the program, follow the steps:

1. Place your data file into the directory of the project. Alternatively, you can specify the path to the file while running the program.

2. Run the command `go run .` to execute the program.

3. If the file is in the project directory under the name `loras.txt`, the program will use it automatically, else you will be prompted to enter the filename.
   ### Example file
   ```text
   ____________SAMPLERS____________
   euler;
   euler_ancestral;
   uni_pc_bh2
      
   ___________SCHEDULERS___________
   normal;
   karras;
   exponential;
   sgm_uniform;
   simple;
   ddim_uniform
      
   _____________VAES_______________
   kl-f8-anime2.ckpt;
   vae-ft-mse-840000-ema-pruned.ckpt
      
   ___________CHECKPOINTS__________
   modelnamewithoutfolder.ckpt;
   anime\modelname.safetensors;
   kl-f8-anime2.vae.ckpt;
   sdxl\768-v-ema.safetensors
      
   _____________LORAS______________
   lorafilenameNoFolder.safetensors;
   artist\artistname.safetensors;
   character\charactername.safetensors
   ```

4. After the successful execution of the program, a JSON file named `models.json` containing the categorized data will be generated in the project directory.

Please ensure you have Go SDK 1.21.1 or later installed as the project uses Go programming language version 1.21.