savedcmd_/home/erazoex/Repositories/SO1_1S2024_201807253/Proyecto1/ram/ram.ko := ld -r -m elf_x86_64 -z noexecstack --build-id=sha1  -T scripts/module.lds -o /home/erazoex/Repositories/SO1_1S2024_201807253/Proyecto1/ram/ram.ko /home/erazoex/Repositories/SO1_1S2024_201807253/Proyecto1/ram/ram.o /home/erazoex/Repositories/SO1_1S2024_201807253/Proyecto1/ram/ram.mod.o;  make -f ./arch/x86/Makefile.postlink /home/erazoex/Repositories/SO1_1S2024_201807253/Proyecto1/ram/ram.ko