import { clearToken } from "@/api/mande";
import { User, Home } from "@/icons";
import { Factory } from "@/icons";
import { useAppState } from "@/store";
import { removeToken } from "@/utils";
import { useNavigate, A } from "@solidjs/router";
import type { ParentProps } from "solid-js";
import { Popover, PopoverButton, PopoverPanel, Button } from "terracotta";

export default function Layout(props: ParentProps) {
	const navigate = useNavigate();
	const { clearAppState } = useAppState();

	const logout = () => {
		removeToken();
		clearToken();
		clearAppState();
		navigate("/auth/signin", { replace: true });
	};

	console.log(logout);

	return (
		<div class="w-[1000px] max-w-full px-4 mx-auto">
			<div class="navbar">
				<div class="flex-1">
					<a class="btn btn-ghost text-xl" href="/">OLI AI</a>
				</div>

				<div class="flex-none gap-2">
					<Popover defaultOpen={false} class="relative">
						{({ isOpen }) => (
							<>
								<PopoverButton>
									<span class="btn btn-ghost btn-circle" classList={{ "btn-active": isOpen() }}>
										<User class="text-2xl" />
									</span>
								</PopoverButton>
								<PopoverPanel class="z-10 p-2 shadow absolute end-0 min-w-[200px]">
									<div class="menu bg-base-200 rounded-box">
										<Button as="button" onClick={logout} class="btn btn-sm justify-start btn-ghost">
											Cerrar Sesión
										</Button>
									</div>
								</PopoverPanel>
							</>
						)}
					</Popover>
				</div>
			</div>
			<main class="flex">
				<nav class="w-[200px]">
					<div class="menu w-full rounded-box text-base font-semibold gap-1">
						<Button as={A} href="/home" activeClass="btn-active" class="btn justify-start btn-ghost">
							<Home class="text-2xl" /> Inicio
						</Button>
						<Button as={A} href="/companies" activeClass="btn-active" class="btn justify-start btn-ghost">
							<Factory class="text-2xl" /> Empresas
						</Button>
					</div>
				</nav>
				<div class="flex-1 pl-4">
					{props.children}
				</div>
			</main>
		</div>
	);
}
